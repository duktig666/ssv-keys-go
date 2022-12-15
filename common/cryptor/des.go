// description: des
// @author renshiwei
// Date: 2022/11/3 17:34

package cryptor

import (
	"crypto/cipher"
	"crypto/des"
)

// DesECBEncrypt
//---------------DES ECB加密--------------------
// data: 明文数据
// key: 密钥字符串
// 返回密文数据
func DesECBEncrypt(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	// pkcs5填充
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

// DesECBDecrypter
//---------------DES ECB解密--------------------
// data: 密文数据
// key: 密钥字符串
// 返回明文数据
func DesECBDecrypter(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	// pkcs5填充
	out = pkcs5UnPadding(out)

	return out
}

// DesCBCEncrypt
//---------------DES CBC加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCBCEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText
}

// DesCBCDecrypter
//---------------DES CBC解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCBCDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

// DesCTREncrypt
//---------------DES CTR加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCTREncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCTR(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesCTRDecrypter
//---------------DES CTR解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCTRDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCTR(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

// DesOFBEncrypt
//---------------DES OFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesOFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewOFB(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesOFBDecrypter
//---------------DES OFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesOFBDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewOFB(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}

// DesCFBEncrypt
//---------------DES CFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCFBDecrypter(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesCFBDecrypter
//---------------DES CFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCFBDecrypter(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCFBEncrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = pkcs5UnPadding(cryptText)

	return cryptText
}
