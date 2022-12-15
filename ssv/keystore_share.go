// description: ETH keystore share implement
// @author renshiwei
// Date: 2022/11/1 19:33

package ssv

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/duktig666/ssv-keys-go/common/cryptor"
	"github.com/duktig666/ssv-keys-go/eth1"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

type IShare struct {
	privateKey string
	publicKey  string
	id         interface{}
}

// Payload ssv 智能合约交易信息
type Payload struct {
	ValidatorPublicKey  string   `json:"validatorPublicKey"`
	OperatorIds         string   `json:"operatorIds"`
	SharePublicKeys     []string `json:"sharePublicKeys"`
	encryptedKeys       []string
	AbiSharePrivateKeys []string `json:"sharePrivateKey"`
	SsvAmount           string   `json:"ssvAmount"`
}

// Operator ssv operator 信息
type Operator struct {
	Id        int    `json:"id"`
	PublicKey string `json:"publicKey"`
	Fee       uint64 `json:"-"`
}

// Shares keystore 分片信息
type Shares struct {
	PublicKeys    []string `json:"publicKeys"`
	EncryptedKeys []string `json:"encryptedKeys"`
}

type KeystoreShareData struct {
	PublicKey string      `json:"publicKey"`
	Operators []*Operator `json:"operators"`
	Shares    *Shares     `json:"shares"`
}

// KeystoreShareRes keystore分片返回的最终结果
type KeystoreShareRes struct {
	Version string `json:"version"`

	Data *KeystoreShareData `json:"data"`

	Payload struct {
		Readable *Payload `json:"readable"`
		Raw      string   `json:"raw"`
	} `json:"payload"`

	CreatedAt time.Time `json:"createdAt"`
}

// KeystoreShareInfo ssv KeystoreShare 业务计算结构
type KeystoreShareInfo struct {
	// 分片的索引号
	ID uint64
	// 分片公钥
	PublicKey string
	// 分片私钥（受限访问）
	secretKey string
	// 分片加密后的私钥
	EncryptedKey    string
	AbiEncryptedKey string
	// 该分片选择的operator
	Operator *Operator
}

// CreateThreshold receives a bls.SecretKey hex and count.
// Will split the secret key into count shares
func CreateThreshold(skBytes []byte, operators []*Operator) (map[uint64]*IShare, error) {
	threshold := uint64(len(operators))

	// master key Polynomial
	msk := make([]bls.SecretKey, threshold)
	mpk := make([]bls.PublicKey, threshold)

	sk := &bls.SecretKey{}
	if err := sk.Deserialize(skBytes); err != nil {
		return nil, err
	}
	msk[0] = *sk
	mpk[0] = *sk.GetPublicKey()

	_F := (threshold - 1) / 3

	// Receives list of operators IDs. len(operator IDs) := 3 * F + 1
	// construct poly
	for i := uint64(1); i < threshold-_F; i++ {
		sk := bls.SecretKey{}
		sk.SetByCSPRNG()
		msk[i] = sk
		mpk[i] = *sk.GetPublicKey()
	}

	// evaluate shares - starting from 1 because 0 is master key
	shares := make(map[uint64]*IShare)
	for i := uint64(1); i <= threshold; i++ {
		blsID := bls.ID{}

		// blsID 设置为operatorId （golang和ts实现的不同之处）
		operatorId := operators[i-1].Id

		err := blsID.SetDecString(fmt.Sprintf("%d", operatorId))
		if err != nil {
			return nil, err
		}

		sk := bls.SecretKey{}

		err = sk.Set(msk, &blsID)
		if err != nil {
			return nil, err
		}

		pk := bls.PublicKey{}
		err = pk.Set(mpk, &blsID)
		if err != nil {
			return nil, err
		}

		shares[i] = &IShare{
			privateKey: sk.SerializeToHexStr(),
			publicKey:  pk.SerializeToHexStr(),
			id:         blsID.GetHexString(),
		}
	}
	return shares, nil
}

// EncryptShares 构造分片的加密公私钥对
func EncryptShares(skBytes []byte, operators []*Operator) ([]*KeystoreShareInfo, error) {
	shareCount := uint64(len(operators))
	shares, err := CreateThreshold(skBytes, operators)
	if err != nil {
		return nil, errors.Wrap(err, "creating threshold err")
	}

	keystoreShareInfos := make([]*KeystoreShareInfo, 0, shareCount)

	// 遍历 构造分片rsa & abi加密数据

	for i := 0; i < int(shareCount); i++ {
		share := shares[uint64(i)+1]

		operator := operators[i]
		opk, err := base64.StdEncoding.DecodeString(operator.PublicKey)
		if err != nil {
			return nil, errors.Wrapf(err, "operator pubkey decode err. pubkey: %s", operator.PublicKey)
		}

		shareSk := "0x" + share.privateKey
		sharePk := "0x" + share.publicKey

		decryptShareSecret, err := cryptor.PublicEncrypt(shareSk, string(opk))
		abiShareSecret, err := eth1.AbiCoder([]string{"string"}, []interface{}{decryptShareSecret})

		keystoreShareInfo := &KeystoreShareInfo{
			ID:              uint64(i),
			PublicKey:       sharePk,
			secretKey:       shareSk,
			EncryptedKey:    decryptShareSecret,
			AbiEncryptedKey: "0x" + hex.EncodeToString(abiShareSecret),
			Operator:        operator,
		}
		keystoreShareInfos = append(keystoreShareInfos, keystoreShareInfo)
	}
	return keystoreShareInfos, nil
}

// 构建 payload 结构体
func buildPayload(validatorPublicKey, ssvAmount string, operators []*Operator, keystoreShareHelpers []*KeystoreShareInfo) *Payload {
	operatorIds := ""
	for _, operator := range operators {
		operatorIds += strconv.Itoa(operator.Id) + ","
	}
	operatorIds = strings.TrimRight(operatorIds, ",")

	count := len(keystoreShareHelpers)
	sharePublicKeys := make([]string, 0, count)
	abiSharePrivateKeys := make([]string, 0, count)
	encryptedSharePrivateKeys := make([]string, 0, count)

	for _, helper := range keystoreShareHelpers {
		sharePublicKeys = append(sharePublicKeys, helper.PublicKey)
		abiSharePrivateKeys = append(abiSharePrivateKeys, helper.AbiEncryptedKey)
		encryptedSharePrivateKeys = append(encryptedSharePrivateKeys, helper.EncryptedKey)
	}

	payload := &Payload{
		ValidatorPublicKey:  validatorPublicKey,
		OperatorIds:         operatorIds,
		SharePublicKeys:     sharePublicKeys,
		AbiSharePrivateKeys: abiSharePrivateKeys,
		SsvAmount:           ssvAmount,
		encryptedKeys:       encryptedSharePrivateKeys,
	}

	return payload
}

func buildPayloadRaw(payload *Payload) string {
	sharePubkeysStr := strings.Replace(strings.Trim(fmt.Sprint(payload.SharePublicKeys), "[]"), " ", ",", -1)
	abiShareSecretsStr := strings.Replace(strings.Trim(fmt.Sprint(payload.AbiSharePrivateKeys), "[]"), " ", ",", -1)

	raw := fmt.Sprintf("%s,%s,%s,%s,%s", payload.ValidatorPublicKey, payload.OperatorIds, sharePubkeysStr, abiShareSecretsStr, payload.SsvAmount)
	return raw
}

func buildKeystoreShareRes(validatorPublicKey, version, ssvAmount string, operators []*Operator, keystoreShareHelpers []*KeystoreShareInfo) *KeystoreShareRes {
	payload := buildPayload(validatorPublicKey, ssvAmount, operators, keystoreShareHelpers)
	raw := buildPayloadRaw(payload)

	keystoreShareData := &KeystoreShareData{
		PublicKey: validatorPublicKey,
		Operators: operators,
		Shares: &Shares{
			PublicKeys:    payload.SharePublicKeys,
			EncryptedKeys: payload.encryptedKeys,
		},
	}

	keystoreShareRes := &KeystoreShareRes{
		Version:   version,
		CreatedAt: time.Now(),
		Data:      keystoreShareData,
	}
	keystoreShareRes.Payload.Readable = payload
	keystoreShareRes.Payload.Raw = raw

	return keystoreShareRes
}

// KeystoreShareV2 keystore 分片结果
func KeystoreShareV2(validatorPublicKey, version, ssvAmount string, skBytes []byte, operators []*Operator) (*KeystoreShareRes, error) {
	keystoreShareInfos, err := EncryptShares(skBytes, operators)
	if err != nil {
		return nil, errors.Unwrap(err)
	}
	keystoreSharesRes := buildKeystoreShareRes(validatorPublicKey, version, ssvAmount, operators, keystoreShareInfos)

	return keystoreSharesRes, nil
}

// KeystoreShareV2ForJson keystore 分片json结果
func KeystoreShareV2ForJson(validatorPublicKey, version, ssvAmount string, skBytes []byte, operators []*Operator) (string, error) {
	keystoreSharesRes, err := KeystoreShareV2(validatorPublicKey, version, ssvAmount, skBytes, operators)
	if err != nil {
		return "", errors.Unwrap(err)
	}

	jsonRes, err := json.Marshal(keystoreSharesRes)
	if err != nil {
		return "", errors.Wrap(err, "failed to json marshal.")
	}

	return string(jsonRes), nil
}
