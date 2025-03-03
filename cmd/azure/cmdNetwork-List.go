package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ipAddresses bool
)

// networkCmd represents the network command
var networkListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("network called")
	},
}

func init() {
	networkCmd.AddCommand(networkListCmd)
	networkListCmd.Flags().BoolVarP(&ipAddresses, "ipAddresses", "i", false, "List IP Addresses for selected tenant")
}
