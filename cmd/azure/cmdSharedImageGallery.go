/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	galleryName string
)

// networkCmd represents the network command
var sharedImageGalleryCmd = &cobra.Command{
	Use:     "sharedImageGallery",
	Aliases: []string{"sig"},
	Short:   "Network related commands",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sharedImageGallery called")

	},
}

func init() {
	azCmd.AddCommand(sharedImageGalleryCmd)

	sharedImageGalleryCmd.PersistentFlags().StringVarP(&galleryName, "galleryName", "n", "", "Shared Image Gallery name")
	sharedImageGalleryCmd.MarkPersistentFlagRequired("galleryName")
}
