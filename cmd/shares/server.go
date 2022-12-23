// description:
// @author renshiwei
// Date: 2022/12/15 14:48

package shares

import (
	"fmt"
	"github.com/duktig666/ssv-keys-go/common/global"
	"github.com/duktig666/ssv-keys-go/common/initialize"
	"github.com/duktig666/ssv-keys-go/common/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	keystorePath       string
	password           string
	operatorPubkeyList []string
	operatorIdList     []uint32
	ssvAmount          string
	outputPath         string
	shareCount         uint32
)

var (
	StartCmd = &cobra.Command{
		Use:     "shares",
		Short:   "ssv keystore shares",
		Example: global.Config.Cli.Name + ` shares --keystore ./temp/input/keystore-xxx.json --password <your password> --operators ...,...,...,... operator-ids ...,...,...,... --output ./temp/output/shares.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(color.GreenString("set operators and operator-ids. count:%v\n", shareCount))
	operatorIdList = make([]uint32, 0, shareCount)
	operatorPubkeyList = make([]string, 0, shareCount)

	for i := uint32(0); i < shareCount; i++ {
		var operatorId uint32
		var operatorPubkey string

		fmt.Println(color.GreenString("Set %v Operator id", i+1))

		_, err := fmt.Scan(&operatorId)
		if err != nil {
			fmt.Println(color.RedString("Set %v Operator id err:%v\n", i+1, err))
		}

		fmt.Println(color.GreenString("Set %v Operator key", i+1))
		_, err = fmt.Scan(&operatorPubkey)
		if err != nil {
			fmt.Println(color.RedString("Set %v Operator key err:%v\n", i+1, err))
		}

		operatorIdList = append(operatorIdList, operatorId)
		operatorPubkeyList = append(operatorPubkeyList, operatorPubkey)
	}

	shareRes, err := keystoreShare()
	if err != nil {
		fmt.Println(color.RedString("shares fail. err:%v", err))
	}

	if outputPath == "" {
		fmt.Println(color.GreenString("shares res"))
		fmt.Println(shareRes)
	} else {
		err = output([]byte(shareRes))
		if err != nil {
			fmt.Println(color.RedString("shares fail. err:%v", err))
		}
	}

	return nil
}

func init() {
	err := initialize.InitBls()
	if err != nil {
		logger.Errorf("init BLS error: %v", err)
	}
	logger.InitLog()

	// add flags
	StartCmd.Flags().StringVarP(&keystorePath, "keystore", "k", "", "keystore.json file path")
	StartCmd.Flags().StringVarP(&password, "password", "p", "", "keystore password")
	StartCmd.Flags().StringVarP(&ssvAmount, "ssv-amount", "s", "", "SSV Token amount fee required for this transaction in Wei. (default:'0')")
	StartCmd.Flags().StringVar(&outputPath, "output", "", "Result output path: (eg: ./shares.json). If not set, output to console.")
	StartCmd.Flags().Uint32Var(&shareCount, "count", 4, "share count (default:4)")

	// operators
	//StartCmd.Flags().StringVarP(&operatorPubkeys, "operators", "o", "", "Comma-separated list of the operator keys")
	//StartCmd.Flags().StringVarP(&operatorIds, "operator-ids", "i", "", "Comma-separated list of the operator ids (same sequence as operators)")

	// set flags relation
	err = StartCmd.MarkFlagRequired("keystore")
	if err != nil {
		logger.Errorf("keystore is required. error: %v", err)
	}

	err = StartCmd.MarkFlagRequired("password")
	if err != nil {
		logger.Errorf("password is required. error: %v", err)
	}
	StartCmd.MarkFlagsRequiredTogether("keystore", "password")

	//err = StartCmd.MarkFlagRequired("operators")
	//if err != nil {
	//	logger.Errorf("operators is required. error: %v", err)
	//}
	//
	//err = StartCmd.MarkFlagRequired("operator-ids")
	//if err != nil {
	//	logger.Errorf("operator-ids is required. error: %v", err)
	//}

}
