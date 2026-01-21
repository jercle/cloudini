package m365

import (
	"github.com/spf13/cobra"
)

// subsCmd represents the subs command
var sharepointCmd = &cobra.Command{
	Use:     "sharepoint",
	Aliases: []string{"sp"},
	Short:   "Sharepoint Online",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("subs called")
		// azProfile, _ := getSubs()
		// fmt.Println("print")

		// azProfile.PrintSubs()
		// azProfile.Sort()

	},
}

func init() {
	m365Cmd.AddCommand(sharepointCmd)
	// subsCmd.Flags().BoolVarP(&SetActive, "setActive", "x", false, "Change active Azure subscription")
	// subsCmd.Flags().BoolVarP(&ShowActive, "showActive", "a", false, "Show current active Azure subscription")
	// subsCmd.Flags().BoolVarP(&Fetch, "fetch", "f", false, "Fetch all available subscriptions from Azure")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
