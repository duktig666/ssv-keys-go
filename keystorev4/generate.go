// description:
// @author renshiwei
// Date: 2022/8/18 17:43

package keystorev4

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/duktig666/ssv-keys-go/common/cryptor"
	"github.com/google/uuid"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

// GenerateKeystorev4AndPass @return Keystorev4;pass;err
func GenerateKeystorev4AndPass() (*Keystorev4, string, *e2types.BLSPrivateKey, error) {
	// 生成密码
	pass, err := generatePass()
	if err != nil {
		return nil, "", nil, errors.New("generate pass error")
	}

	blsKey, err := GenerateRandomBLSPrivateKey()
	if err != nil {
		return nil, "", nil, errors.New("GenerateBLSPrivateKey error")
	}

	keystore, err := GenerateKeystoreV4Custom(blsKey, pass)
	if err != nil {
		return nil, "", nil, err
	}

	return keystore, pass, blsKey, nil
}

//GenerateKeystoreV4JsonAndPass
//@return keystoreV4;pass;pubkey;BLSPrivateKey;err
func GenerateKeystoreV4JsonAndPass() (string, string, string, *e2types.BLSPrivateKey, error) {
	keystore, pass, blsKey, err := GenerateKeystorev4AndPass()
	if err != nil {
		return "", "", "", nil, err
	}

	keystoreByte, err := json.Marshal(keystore)
	if err != nil {
		return "", "", "", nil, errors.New("keystoreV4 to json error")
	}

	return string(keystoreByte), pass, keystore.Pubkey, blsKey, nil
}

//GenerateRandomBLSPrivateKey 生成BLSPrivateKey
func GenerateRandomBLSPrivateKey() (*e2types.BLSPrivateKey, error) {
	blsPrivateKey, err := e2types.GenerateBLSPrivateKey()
	if err != nil {
		return nil, err
	}
	return blsPrivateKey, nil
}

//GenerateKeystoreV4Custom 生成 keystoreV4的struct
func GenerateKeystoreV4Custom(blsPrivateKey *e2types.BLSPrivateKey, pass string) (*Keystorev4, error) {
	pubkey := hex.EncodeToString(blsPrivateKey.PublicKey().Marshal())

	// 私钥
	secret := blsPrivateKey.Marshal()
	keystoreMap, err := Pbkdf2Encryptor.Encrypt(secret, pass)
	if err != nil {
		return nil, errors.New("encrypt BLSPrivateKey error")
	}

	// keystoreMap中解析成 keystore.json
	marshal, err := json.Marshal(keystoreMap)
	if err != nil {
		return nil, errors.New("privateKey to json error")
	}

	var keystore Keystorev4

	// keystore Crypto
	err = json.Unmarshal(marshal, &keystore.Crypto)
	if err != nil {
		return nil, errors.New("privateKey json to Keystorev4 struct error")
	}

	// uuid
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("uuid error")
	}

	// 设置keystore的其他属性
	keystore.Version = 4
	keystore.Uuid = newUUID.String()
	keystore.Pubkey = pubkey

	return &keystore, nil
}

// 随机生成安全的密码
func generatePass() (string, error) {
	return cryptor.GenerateRandomPass(32, 10, 10, false, false)
}
