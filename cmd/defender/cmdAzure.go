package defender

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
	useAzCliAuth   bool
)

var defenderCmd = &cobra.Command{
	Use:     "defender",
	Aliases: []string{"xdr"},
	Short:   "Microsoft Defender",
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
	cmd.RootCmd.AddCommand(defenderCmd)
	defenderCmd.PersistentFlags().StringVar(&clientId, "clientId", "", "Client ID for Service Principal authentication.")
	defenderCmd.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Client Secret for Service Principal authentication.")
	defenderCmd.PersistentFlags().StringVar(&tenantId, "tenantId", "", "Tenant ID.")
	defenderCmd.PersistentFlags().BoolVar(&useAzCliAuth, "useAzCliAuth", false, "Uses Azure CLI authentication instead of a Service Principal")

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
	// defenderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defenderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
