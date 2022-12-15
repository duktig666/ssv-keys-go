// description:
// @author renshiwei
// Date: 2022/12/15 14:46

package initialize

import e2types "github.com/wealdtech/go-eth2-types/v2"

func InitBls() error {
	// 执行bls的相关方法必须初始化
	if err := e2types.InitBLS(); err != nil {
		return err
	}
	return nil
}
