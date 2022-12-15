// description:
// @author renshiwei
// Date: 2022/8/18 19:15

package cryptor

import "github.com/sethvargo/go-password/password"

// GenerateRandomPass 随机生成安全的密码
func GenerateRandomPass(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error) {
	return password.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
}
