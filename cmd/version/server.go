package version

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const Version = "0.0.1"

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Get version info",
		Example: "ssv-key version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf("ssv-key version: %s\n", color.GreenString(Version))
	return nil
}
