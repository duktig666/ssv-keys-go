// description:
// @author renshiwei
// Date: 2022/9/14 17:08

package ssv

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/pkg/errors"
)

var keySize = 2048

//GenerateKeysBase64 ssv 公私钥生成
func GenerateKeysBase64() (string, string, error) {
	pkb, skb, err := GenerateKeys()
	return base64.StdEncoding.EncodeToString(pkb), base64.StdEncoding.EncodeToString(skb), err
}

// GenerateKeys using rsa random generate keys and return []byte bas64
func GenerateKeys() ([]byte, []byte, error) {
	// generate random private key (secret)
	sk, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to generate rsa key")
	}
	// retrieve public key from the newly generated secret
	pk := &sk.PublicKey

	// convert to bytes
	skPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(sk),
		},
	)
	pkBytes, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to marshal public key")
	}
	pkPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pkBytes,
		},
	)
	return pkPem, skPem, nil
}

// DecodeKey with secret key (base64) and hash (base64), return the encrypted key string
func DecodeKey(sk *rsa.PrivateKey, hashBase64 string) (string, error) {
	hash, _ := base64.StdEncoding.DecodeString(hashBase64)
	decryptedKey, err := rsa.DecryptPKCS1v15(rand.Reader, sk, hash)
	if err != nil {
		return "", errors.Wrap(err, "Failed to decrypt key")
	}
	return string(decryptedKey), nil
}

// ConvertPemToPrivateKey return rsa private key from secret key
func ConvertPemToPrivateKey(skPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(skPem))
	// TODO: resolve deprecation https://github.com/golang/go/issues/8860
	enc := x509.IsEncryptedPEMBlock(block) //nolint
	b := block.Bytes
	if enc {
		var err error
		// TODO: resolve deprecation https://github.com/golang/go/issues/8860
		b, err = x509.DecryptPEMBlock(block, nil) //nolint
		if err != nil {
			return nil, errors.Wrap(err, "Failed to decrypt private key")
		}
	}
	parsedSk, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse private key")
	}
	return parsedSk, nil
}

// PrivateKeyToByte converts privateKey to []byte
func PrivateKeyToByte(sk *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(sk),
		},
	)
}

// ExtractPublicKey get public key from private key and return base64 encoded public key
func ExtractPublicKey(sk *rsa.PrivateKey) (string, error) {
	pkBytes, err := x509.MarshalPKIXPublicKey(&sk.PublicKey)
	if err != nil {
		return "", errors.Wrap(err, "Failed to marshal private key")
	}
	pemByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pkBytes,
		},
	)

	return base64.StdEncoding.EncodeToString(pemByte), nil
}
