package azure

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/spf13/cobra"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

// var SetActive bool = false
var filterName string
var filterByRetentionInDaysAsDefault string
var caseInsensitive bool
var updateLaTables bool

// subsCmd represents the subs command
var laTablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "Azure Log Analytics Tables",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		// resourceGroup := cmd.Flag("resourceGroup").Value.String()

		if resourceGroup == "" {
			// cmd.PrintErrln("test")
			fmt.Println("--resourceGroup[-r] must be supplied\n")
			cmd.Help()
			os.Exit(1)
		}

		outJSON, err := cmd.Flags().GetBool("outJSON")
		if err != nil {
			log.Fatal(err)
		}

		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatal(err)
		}

		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Fetching workspace tables")
		w.Start()
		data, err := getAllWorkspaceTables(cred, subscriptionId, resourceGroup, workspaceName)
		// fmt.Println(data)
		w.Stop()
		w.PersistWith(spin.Spinner{Frames: []string{"👍 "}}, fmt.Sprint(len(data), " tables fetched"))
		if err != nil {
			log.Fatal(err)
		}

		if filterName != "" {
			data.filterByName(filterName, caseInsensitive)
		}

		if filterByRetentionInDaysAsDefault != "" {
			if filterByRetentionInDaysAsDefault == "true" {
				data.filterByRetentionInDaysAsDefault(true)
			} else if filterByRetentionInDaysAsDefault == "false" {
				data.filterByRetentionInDaysAsDefault(false)
			}
		}

		if filterName != "" || filterByRetentionInDaysAsDefault != "" {
			fmt.Println("Results with filters applied: ", len(data))
		}

		if updateLaTables {
			fmt.Println("Update called")
		}

		// outJSON := cmd.Flag("outJSON").Value

		// fmt.Println()
		if !outJSON {
			data.printTable()
		} else {
			data.printJSON()
		}
	},
}

func init() {
	// log.Fatal(cmd.root.PersistentFlags().Lookup("outJSON").Value.Type())
	logAnalyticsCmd.AddCommand(laTablesCmd)

	// if laTablesCmd.CommandPath() == "la tables" {
	// 	cobra.MarkFlagRequired(azCmd.PersistentFlags(), "resourceGroup")
	// }
	laTablesCmd.PersistentFlags().StringVarP(&filterName, "filterByName", "n", "", "Filter results by name")
	laTablesCmd.PersistentFlags().BoolVarP(&caseInsensitive, "caseInsensitive", "i", false, "Case insensitive search")
	laTablesCmd.PersistentFlags().StringVarP(&filterByRetentionInDaysAsDefault, "filterByRetentionInDaysAsDefault", "d", "", "Filter results by filterByRetentionInDaysAsDefault property")
	laTablesCmd.PersistentFlags().BoolVar(&updateLaTables, "updateTables", false, "Switch to update tables instead of just listing")
}
