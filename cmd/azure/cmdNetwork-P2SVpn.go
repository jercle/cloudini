package azure

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	storageAccountName      string
	containerName           string
	p2sVpnGatewayResourceId string
)

// subsCmd represents the subs command
var networkP2SVpnCmd = &cobra.Command{
	Use: "p2svpn",
	// Aliases: []string{"p2svpn"},
	Short: "Azure Point-to-Site VPN commands",
	Run: func(cmd *cobra.Command, args []string) {
		if tenantName == "" || storageAccountName == "" || containerName == "" || p2sVpnGatewayResourceId == "" {
			fmt.Println("Please provide tenantName, storageAccountName, containerName, and p2sVpnGatewayResourceId")
			os.Exit(1337)
		}
		GenerateP2SVpnConnectionHealthDetailed(p2sVpnGatewayResourceId, tenantName, storageAccountName, containerName)
	},
}

func init() {
	networkCmd.AddCommand(networkP2SVpnCmd)
	networkP2SVpnCmd.Flags().StringVarP(&tenantName, "tenantName", "t", "", "Name of Azure Tenant")
	networkP2SVpnCmd.Flags().StringVarP(&storageAccountName, "storageAccountName", "s", "", "Name of Storage Account to upload P2S VPN connection statistics")
	networkP2SVpnCmd.Flags().StringVarP(&containerName, "containerName", "c", "", "Name of Blob Container to upload P2S VPN connection statistics")
	networkP2SVpnCmd.Flags().StringVarP(&p2sVpnGatewayResourceId, "p2sVpnGatewayResourceId", "p", "", "Resource ID of P2S VPN Gateway")

	// networkP2SVpnCmd.Flags().StringVarP(&galleryImageName, "imageName", "i", "", "Compute Gallery Image name")
	// networkP2SVpnCmd.MarkFlagRequired("galleryImageName")
	// networkP2SVpnCmd.Flags().BoolVarP(&getLatestVersionNumber, "getLatestVersionNumber", "l", false, "Get latest version number")
	// networkP2SVpnCmd.Flags().BoolVarP(&getNewVersionPatchNumber, "getNewVersionPatchNumber", "p", false, "Increment version patch number")
	// networkP2SVpnCmd.MarkFlagsMutuallyExclusive("getLatestVersionNumber", "getNewVersionPatchNumber")
	// networkP2SVpnCmd.Flags().StringVarP(&checkVersionExists, "checkVersionExists", "c", "", "Compute Gallery Image version")

}
