package defender

import (
	"fmt"

	"github.com/spf13/cobra"
)

var machinesCmd = &cobra.Command{
	Use:     "machines",
	Aliases: []string{"m"},
	Short:   "Machines/Devices in Defender XDR",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("network called")
	},
}

func init() {
	defenderCmd.AddCommand(machinesCmd)
	// machinesCmd.Flags().
}
