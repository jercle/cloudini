package azure

import (
	"fmt"
	"os"
	"strings"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var galleryImageName string
var checkVersionExists string
var getLatestVersionNumber bool
var getNewVersionPatchNumber bool

// subsCmd represents the subs command
var cmdComputeGalleryImage = &cobra.Command{
	Use: "image",
	// Aliases: []string{"image"},
	Short: "Azure Compute Gallery / Shared Image Gallery Images",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			authOpts lib.AzureMultiAuthTokenRequestOptions
			token    *lib.AzureMultiAuthToken
			err      error
		)

		if clientId != "" && clientSecret != "" && tenantId != "" {
			authOpts.ClientID = clientId
			authOpts.ClientSecret = clientSecret
			authOpts.TenantID = tenantId
			// authOpts.TenantName = tenantName

			token, err = GetServicePrincipalMultiAuthToken(authOpts)
			lib.CheckFatalError(err)

			imageDefinition := GetGalleryImage(subscriptionId, resourceGroup, galleryName, galleryImageName, token)
			_ = imageDefinition

			if getLatestVersionNumber {

				versions, _ := GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
				_, latest := versions.Latest()

				fmt.Println(latest)
			}

			if getNewVersionPatchNumber {
				versions, _ := GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
				if len(versions.Versions) == 0 {
					fmt.Println("")
				} else {
					latest, _ := versions.Latest()
					fmt.Println(latest.IncrementPatchVersion())
				}
			}

			if checkVersionExists != "" {
				versions, _ := GetGalleryImageVersions(subscriptionId, resourceGroup, galleryName, galleryImageName, *token)
				versionExists := versions.CheckVersionExists(checkVersionExists)

				if versionExists {
					fmt.Println("Version " + checkVersionExists + " exists")
					os.Exit(0)
				} else {
					sortedVersions := versions.ListVersions()
					err := fmt.Errorf("Version " + checkVersionExists + " does not exist. Available versions: " + strings.Join(sortedVersions, ", "))
					lib.CheckFatalError(err)
				}
			}

		} else {
			fmt.Println("Please provide --clientId, --clientSecret, and --tenantId")
		}

		// fmt.Println(authOpts)
		// os.Exit(0)

	},
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
	cmdComputeGallery.AddCommand(cmdComputeGalleryImage)
	cmdComputeGalleryImage.Flags().StringVarP(&galleryImageName, "imageName", "i", "", "Compute Gallery Image name")
	cmdComputeGalleryImage.MarkFlagRequired("galleryImageName")
	cmdComputeGalleryImage.Flags().BoolVarP(&getLatestVersionNumber, "getLatestVersionNumber", "l", false, "Get latest version number")
	cmdComputeGalleryImage.Flags().BoolVarP(&getNewVersionPatchNumber, "getNewVersionPatchNumber", "p", false, "Increment version patch number")
	cmdComputeGalleryImage.MarkFlagsMutuallyExclusive("getLatestVersionNumber", "getNewVersionPatchNumber")
	cmdComputeGalleryImage.Flags().StringVarP(&checkVersionExists, "checkVersionExists", "c", "", "Compute Gallery Image version")

}
