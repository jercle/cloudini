package azure

import (
	"github.com/jercle/cloudini/cmd"
	"github.com/spf13/cobra"
)

var (
	tenantId       string
	subscriptionId string
	resourceGroup  string
	clientSecret   string
	clientId       string
)

var azCmd = &cobra.Command{
	Use:     "azure",
	Aliases: []string{"az"},
	Short:   "Azure",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("az called")
	//	},
}

func init() {
	cmd.RootCmd.AddCommand(azCmd)
	azCmd.PersistentFlags().StringVar(&subscriptionId, "subscriptionId", "", "Subscription ID to run command against. If not supplied, current default Azure CLI subscription is used.")
	azCmd.PersistentFlags().StringVar(&resourceGroup, "resourceGroup", "", "Resource group to run command against.")
	azCmd.PersistentFlags().StringVar(&clientId, "clientId", "", "Client ID for Service Principal authentication.")
	azCmd.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Client Secret for Service Principal authentication.")
	azCmd.PersistentFlags().StringVar(&tenantId, "tenantId", "", "Tenant ID.")

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
