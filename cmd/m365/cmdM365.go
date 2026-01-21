package m365

import (
	"github.com/jercle/cloudini/cmd"
	"github.com/spf13/cobra"
)

// var (
// 	tenantId       string
// 	subscriptionId string
// 	resourceGroup  string
// 	clientSecret   string
// 	clientId       string
// )

var m365Cmd = &cobra.Command{
	Use:     "m365",
	Aliases: []string{"m"},
	Short:   "Office/Microsoft 365",
	Long:    `These commands are related to automation and administration of Microsoft 365`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("az called")
	// },
}

func init() {
	cmd.RootCmd.AddCommand(m365Cmd)
	// azCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "Subscription ID to run command against. If not supplied, current default Azure CLI subscription is used.")
	// azCmd.PersistentFlags().StringVarP(&resourceGroup, "resourceGroup", "r", "", "Resource group to run command against.")
	// azCmd.PersistentFlags().StringVar(&clientId, "clientId", "", "Client ID for Service Principal authentication.")
	// azCmd.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Client Secret for Service Principal authentication.")
	// azCmd.PersistentFlags().StringVarP(&tenantId, "tenantId", "t", "", "Tenant ID.")

	// if subscriptionId == "" {
	// 	sub, err := GetActiveSub()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	subscriptionId = sub.ID
	// }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
