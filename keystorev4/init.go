// description:
// @author renshiwei
// Date: 2022/8/18 18:34

package keystorev4

import (
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

// keystorev4 Encryptor @see https://www.cnblogs.com/wanghui-garcia/p/10007768.html
var (
	// Pbkdf2Encryptor 默认（消耗资源更少，使用范围更广）
	Pbkdf2Encryptor *keystorev4.Encryptor
	ScryptEncryptor *keystorev4.Encryptor
)

func init() {
	Pbkdf2Encryptor = keystorev4.New()
	ScryptEncryptor = keystorev4.New(keystorev4.WithCipher("scrypt"))
}
