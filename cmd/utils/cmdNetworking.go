package utils

import (
	"github.com/spf13/cobra"
)

// utilsCmd represents the util command
var cmdNetworking = &cobra.Command{
	Use:     "networking",
	Short:   "Networking - Subnet calc, etc",
	Aliases: []string{"nw"},
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// configFile, _, _ := lib.InitConfig(nil)
	// 	// fmt.Println(configFile)
	// },
}

func init() {
	cmdUtils.AddCommand(cmdNetworking)
	// utilsCmd.Flags().IntVarP("pwdgen", )
	// StringVarP("pwdgen", "p", 0)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
