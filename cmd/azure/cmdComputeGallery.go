package azure

import (
	"github.com/spf13/cobra"
)

var galleryName string

// subsCmd represents the subs command
var cmdComputeGallery = &cobra.Command{
	Use:     "computeGallery",
	Aliases: []string{"cg"},
	Short:   "Azure Compute Gallery / Shared Image Gallery",
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
	azCmd.AddCommand(cmdComputeGallery)
	cmdComputeGallery.PersistentFlags().StringVarP(&galleryName, "galleryName", "n", "", "Compute Gallery name")
	cmdComputeGallery.MarkPersistentFlagRequired("galleryName")
}
