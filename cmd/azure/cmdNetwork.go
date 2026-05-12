package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"nw"},
	Short:   "Network related commands",
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
	azCmd.AddCommand(networkCmd)
	// networkCmd.Flags().
}
