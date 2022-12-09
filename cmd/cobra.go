// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/duktig666/ssv-keys-go/cmd/version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "ssv-key",
	Short:        "ETH SSV Key's Golang implementation",
	SilenceUsage: true,
	Long:         `ssv-key:https://github.com/duktig666/ssv-keys-go`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welcome to use ` + `ssv-key.` + ` use ` + `-h` + ` see cli`
	usageStr1 := `You can also refer to the related content of https://github.com/qiaoshurui/couples-subtotal`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
	//rootCmd.AddCommand(api.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
