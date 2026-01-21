package utils

import (
	"fmt"

	dt "github.com/jercle/cloudini/cmd/utils/datatransforms"
	"github.com/spf13/cobra"
)

// csv2jsonCmd represents the csv2json command
var csv2jsonCmd = &cobra.Command{
	Use:     "csv2json",
	Short:   "Converts CSV data to JSON",
	Aliases: []string{"c2j"},
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(dt.CsvFileToJson(fileName))
	},
}

func init() {
	datatransformCmd.AddCommand(csv2jsonCmd)
	// csv2jsonCmd.Flags().StringVarP(&jsonString, "jsonString", "j", "", "Json string to transform")
	// csv2jsonCmd.Flags().StringVarP(&fileName, "filename", "f", "", "Csv file transform")
	// csv2jsonCmd.Flags().StringVarP(&webApi, "webApi", "a", "", "Json http API to get data and parse")
	// csv2jsonCmd.Flags().StringVarP(&outFile, "outFile", "o", "", "File to save output")
	// csv2jsonCmd.Flags().BoolVarP(&stdin, "stdin", "i", false, "Used for piping in with stdin")
	// csv2jsonCmd.Flags().BoolVarP(&getFromClipboard, "getFromClipboard", "c", false, "Get json from clipboard")
	// csv2jsonCmd.Flags().BoolVarP(&saveToClipboard, "saveToClipboard", "p", false, "Save output to clipboard")
	// csv2jsonCmd.MarkFlagsMutuallyExclusive("jsonString", "jsonFile", "stdin", "getFromClipboard")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csv2jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csv2jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
