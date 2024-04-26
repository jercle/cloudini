/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package azure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	latest           bool
	galleryImageName string
	versionOnly      bool
)

// networkCmd represents the network command
var sharedImageGalleryVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Network related commands",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if subscriptionId == "" {
		// 	sub, err := GetActiveSub()
		// 	subscriptionId = sub.ID
		// 	lib.CheckFatalError(err)
		// }
		// if resourceGroup == "" {
		// 	lib.CheckFatalError(fmt.Errorf("Must provide resourceGroup [-n]"))
		// 	os.Exit(1)
		// }
		token, err := GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
		lib.CheckFatalError(err)

		var (
			subscriptionId    = "fdeee0c2-5569-40ea-9ad9-81dd325f6e1e"
			resourceGroupName = "rg-apcdtqdesktop-aib"
			galleryName       = "sigapcdtqdesktopaibimages"
			galleryImageName  = "imgdef-specialised-win10-multi-session-gen2"
		)

		images := GetSIGImageVersions(subscriptionId, resourceGroupName, galleryName, galleryImageName, token)

		if versionOnly {
			fmt.Println(images.Latest().Name)
			os.Exit(0)
		}

		if latest {
			jsonBytes, _ := json.MarshalIndent(images.Latest(), "", "  ")
			fmt.Println(string(jsonBytes))
		} else {
			jsonBytes, _ := json.MarshalIndent(images, "", "  ")
			fmt.Println(string(jsonBytes))
		}

	},
}

func init() {
	sharedImageGalleryCmd.AddCommand(sharedImageGalleryVersionsCmd)
	sharedImageGalleryVersionsCmd.Flags().StringVarP(&galleryImageName, "imageName", "i", "", "Name of SIG Image")
	sharedImageGalleryVersionsCmd.Flags().BoolVarP(&versionOnly, "versionOnly", "v", false, "Only return latest version number")
	sharedImageGalleryVersionsCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Only return latest version")
}
