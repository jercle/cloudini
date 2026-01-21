package azure

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var SetActive bool
var ShowActive bool
var Fetch bool

// subsCmd represents the subs command
var subsCmd = &cobra.Command{
	Use:   "subs",
	Short: "Commands related to Azure Subscriptions, as well as changing current subscription of Azure CLI",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("subs called")
		// azProfile, _ := GetCliSubs()
		// fmt.Println("print")

		// azProfile.PrintSubs()
		// azProfile.Sort()

		if ShowActive {
			activeSub, err := GetActiveCliSub()
			lib.CheckFatalError(err)
			jsonStr, _ := json.MarshalIndent(activeSub, "", "  ")
			fmt.Println(string(jsonStr))
		}

		if SetActive {
			_, subs := PromptSelectTenant()
			ChangeActiveSubscription(subs)
		} else {
		}
	},
}

func init() {
	azCmd.AddCommand(subsCmd)
	subsCmd.Flags().BoolVarP(&SetActive, "setActive", "x", false, "Change active Azure subscription")
	subsCmd.Flags().BoolVarP(&ShowActive, "showActive", "a", false, "Show current active Azure subscription")
	subsCmd.Flags().BoolVarP(&Fetch, "fetch", "f", false, "Fetch all available subscriptions from Azure")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
