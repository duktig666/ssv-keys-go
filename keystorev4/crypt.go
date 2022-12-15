// description:Keystorev4 加解密
// @author renshiwei
// Date: 2022/8/20 22:37

package keystorev4

import (
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

// SwitchEncryptor 选择匹配的 Keystorev4 Encryptor
func SwitchEncryptor(keystore *Keystorev4) *keystorev4.Encryptor {
	var encryptor *keystorev4.Encryptor
	switch keystore.Crypto.KDF.Function {
	case "pbkdf2":
		encryptor = Pbkdf2Encryptor
	case "scrypt":
		encryptor = ScryptEncryptor
	default:
		encryptor = Pbkdf2Encryptor
	}
	return encryptor
}

//Decrypt 解密 Keystorev4
func Decrypt(keystore *Keystorev4, pass string) ([]byte, error) {
	pubkey := keystore.Pubkey
	encryptor := SwitchEncryptor(keystore)

	marshal, err := json.Marshal(keystore.Crypto)
	var cryptoMap map[string]interface{}
	err = json.Unmarshal(marshal, &cryptoMap)
	if err != nil {
		return nil, errors.Errorf("keystore格式不符合 error. pubkey:%v", pubkey)
	}

	deSecret, err := encryptor.Decrypt(cryptoMap, pass)
	if err != nil {
		return nil, errors.Errorf("keystore decrypt error. pubkey:%v", pubkey)
	}

	return deSecret, nil
}

//DecryptFromJson 解密 Keystorev4 json str
func DecryptFromJson(keystoreV4Str string, pass string) ([]byte, error) {
	var keystore *Keystorev4
	err := json.Unmarshal([]byte(keystoreV4Str), &keystore)
	if err != nil {
		return nil, errors.Errorf("keystore json str 格式不符合 error.")
	}
	return Decrypt(keystore, pass)
}

//Encrypt Keystorev4 privateSecret encrypt
func Encrypt(privateSecret, pass string) (*Keystorev4, error) {
	decodeString, err := hex.DecodeString(privateSecret)

	blsPrivateSecret, err := e2types.BLSPrivateKeyFromBytes(decodeString)
	if err != nil {
		return nil, errors.Errorf("privateSecret is involid")
	}
	return GenerateKeystoreV4Custom(blsPrivateSecret, pass)
}

//EncryptForJsonStr Keystorev4 privateSecret encrypt
func EncryptForJsonStr(privateSecret, pass string) (string, error) {
	keystore, err := Encrypt(privateSecret, pass)
	if err != nil {
		return "", err
	}
	keystoreStr, err := json.Marshal(keystore)
	if err != nil {
		return "", errors.Errorf("EncryptForJsonStr json error")
	}
	return string(keystoreStr), nil
}
