package azure

import (
	"github.com/spf13/cobra"
)

// costCmd represents the cost command
var costCmd = &cobra.Command{
	Use:   "cost",
	Short: "Azure Costing",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("cost called")
	//	},
}

func init() {
	azCmd.AddCommand(costCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// costCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// costCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
