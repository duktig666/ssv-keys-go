// description:
// @author renshiwei
// Date: 2022/12/15 14:48

package shares

import (
	"fmt"
	"github.com/duktig666/ssv-keys-go/common/global"
	"github.com/duktig666/ssv-keys-go/common/initialize"
	"github.com/duktig666/ssv-keys-go/common/logger"
	"github.com/spf13/cobra"
)

var (
	keystorePath string
	password     string
	// , 间隔（可切割数组）
	operatorPubkeys string
	// , 间隔（可切割数组）
	operatorIds string
	ssvAmount   string
	outputPath  string
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
	shareRes, err := keystoreShare()
	if err != nil {
		logger.Errorf("shares fail. err:%v", err)
	}

	if outputPath == "" {
		logger.Infof("shares res:\n")
		fmt.Println(shareRes)
	} else {
		err = output([]byte(shareRes))
		if err != nil {
			logger.Errorf("shares fail. err:%v", err)
		}
	}

	return nil
}

func init() {
	err := initialize.InitBls()
	if err != nil {
		logger.Errorf("init BLS error: %v", err)
	}

	// add flags
	StartCmd.Flags().StringVarP(&keystorePath, "keystore", "k", "", "keystore.json file path")
	StartCmd.Flags().StringVarP(&password, "password", "p", "", "keystore password")
	StartCmd.Flags().StringVarP(&operatorPubkeys, "operators", "o", "", "Comma-separated list of the operator keys")
	StartCmd.Flags().StringVarP(&operatorIds, "operator-ids", "i", "", "Comma-separated list of the operator ids (same sequence as operators)")
	StartCmd.Flags().StringVarP(&ssvAmount, "ssv-amount", "s", "", "SSV Token amount fee required for this transaction in Wei. (default:'0')")
	StartCmd.Flags().StringVar(&outputPath, "output", "", "Result output path: (eg: ./shares.json). If not set, output to console.")

	// set flags relation
	err = StartCmd.MarkFlagRequired("keystore")
	if err != nil {
		logger.Errorf("keystore is required. error: %v", err)
	}

	err = StartCmd.MarkFlagRequired("password")
	if err != nil {
		logger.Errorf("password is required. error: %v", err)
	}

	err = StartCmd.MarkFlagRequired("operators")
	if err != nil {
		logger.Errorf("operators is required. error: %v", err)
	}

	err = StartCmd.MarkFlagRequired("operator-ids")
	if err != nil {
		logger.Errorf("operator-ids is required. error: %v", err)
	}

	StartCmd.MarkFlagsRequiredTogether("keystore", "password", "operators", "operator-ids")
}
