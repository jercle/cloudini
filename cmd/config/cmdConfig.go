/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jercle/cloudini/cmd"
	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"

	jsonc "github.com/nwidger/jsoncolor"
)

var SetConfigItem string
var Outfile string
var ExportDecryptedConfig string
var ExportDecryptedTokenCache string
var showConfig bool
var clearTokenCache bool

// Only used when initially encrypting a previously unencrypted config file
var ImportUnencryptedConfigFile string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Cloudini configuration",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("config called")
		// if SetConfigItem != "" {
		// 	strings.Split()
		// }
		// configFile, _, _ := lib.InitConfig(nil)
		// fmt.Println(configFile)

		if showConfig {
			config := lib.GetCldConfig(nil)
			jsonStr, _ := jsonc.MarshalIndent(config, "", "  ")
			fmt.Println(string(jsonStr))
		}

		if ImportUnencryptedConfigFile != "" {
			// fmt.Println(InitialEncryptionOfUnencryptedConfigFile)
			lib.EncryptUnencryptedConfigFile(ImportUnencryptedConfigFile, false)
		}

		configFile, _, cachePath := lib.InitConfig(nil)

		if ExportDecryptedConfig != "" {
			lib.DecryptEncryptedConfigFile(configFile, ExportDecryptedConfig)
		}

		if ExportDecryptedTokenCache != "" {
			fmt.Println(cachePath + "/tkn")
			// os.Exit(0)
			lib.DecryptEncryptedTokenCache(cachePath+"/tkn", ExportDecryptedTokenCache)
		}

		if clearTokenCache {
			lib.ClearTokenCache(nil)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
	// configCmd.Flags().StringVarP(&SetConfigItem, "setItem", "s", "", "Set config item")
	configCmd.Flags().StringVar(&ImportUnencryptedConfigFile, "import", "", "Import unencrypted config file")
	configCmd.Flags().StringVar(&ExportDecryptedConfig, "export", "", "Decrypt and save config file")
	configCmd.Flags().StringVar(&ExportDecryptedTokenCache, "exportToken", "", "Decrypt and save token cache file")
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current config file")
	configCmd.Flags().BoolVarP(&clearTokenCache, "clearTokenCache", "c", false, "Clears cache of token. Use when you get '.../cldFuncs.go:199:0 illegal base64 data at input byte 74480 exit status 1' or similar error")
	// configCmd.Flags().StringVar(&Outfile, "out-file", "", "Output filename")
	// configCmd.MarkFlagsRequiredTogether("export-config", "out-file")
	configCmd.MarkFlagsMutuallyExclusive("import", "export", "show")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
