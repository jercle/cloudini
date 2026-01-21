package ado

import (
	"github.com/jercle/cloudini/cmd"
	"github.com/spf13/cobra"
)

var (
	devOpsOrg           string
	personalAccessToken string
	projectName         string
)

var adoCmd = &cobra.Command{
	Use:   "ado",
	Short: "Azure DevOps related commands",
	// Long: `Commands related to Azure DevOps`,
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

func init() {
	cmd.RootCmd.AddCommand(adoCmd)

	adoCmd.PersistentFlags().StringVarP(&devOpsOrg, "org", "o", "", "Azure DevOps Organization")
	// AZURE_DEVOPS_ORGANIZATION
	adoCmd.PersistentFlags().StringVarP(&personalAccessToken, "pat", "p", "", "Azure DevOps Personal Access Token")
	// AZURE_DEVOPS_EXT_PAT
	adoCmd.PersistentFlags().StringVar(&projectName, "project", "", "Azure DevOps Project")

	adoCmd.MarkPersistentFlagRequired("org")
	adoCmd.MarkPersistentFlagRequired("pat")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
