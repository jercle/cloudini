package web

import (
	"github.com/spf13/cobra"
)

var uri string = ""
var outfile string = ""

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"fetch"},
	Short:   "HTTP GET client",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Get(Request{
			Url:     uri,
			Outfile: outfile,
		})
	},
}

func init() {

	// RootCmd.AddCommand(azure.azCmd)
	// RootCmd.AddCommand(ado.adoCmd)
	// RootCmd.AddCommand(jira.jiraCmd)
	// RootCmd.AddCommand(utils.utilsCmd)

	getCmd.Flags().StringVarP(&uri, "url", "u", "", "Url to get data from (required)")
	getCmd.MarkFlagRequired("url")
	getCmd.Flags().StringVarP(&outfile, "outfile", "o", "", "Filename for output data")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
