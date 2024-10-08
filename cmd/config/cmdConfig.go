/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jercle/cloudini/cmd"
	"github.com/spf13/cobra"
)

var SetConfigItem string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Commands related to Cloudini configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		// if SetConfigItem != "" {
		// 	strings.Split()
		// }
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&SetConfigItem, "setItem", "s", "", "Set config item")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
