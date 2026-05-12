package azure

import (
	"github.com/spf13/cobra"
)

var workspaceName string

// subsCmd represents the subs command
var logAnalyticsCmd = &cobra.Command{
	Use:     "la",
	Aliases: []string{"la"},
	Short:   "Log Analytics / Azure Monitor",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("la called")
	// },
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {

	// },

	// PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {
	// 	// set resourceGroup flag as required for subcommands of this
	// 	azCmd.MarkPersistentFlagRequired("resourceGroup")
	// 	// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
	// 	return cmd.InitializeConfig(ccmd)
	// },
}

func init() {
	azCmd.AddCommand(logAnalyticsCmd)
	logAnalyticsCmd.PersistentFlags().StringVarP(&workspaceName, "workspaceName", "w", "", "LA Workspace name")
	logAnalyticsCmd.MarkPersistentFlagRequired("workspaceName")
}
