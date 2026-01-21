package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ipAddresses bool
)

// networkCmd represents the network command
var networkIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "List IP addresses for selected tenant",
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
	networkCmd.AddCommand(networkIPCmd)
	networkIPCmd.Flags().BoolVarP(&ipAddresses, "ipAddresses", "i", false, "List IP Addresses for selected tenant")
}
