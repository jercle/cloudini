package utils

import (
	"github.com/spf13/cobra"
)

var fileName string
var jsonString string
var webApi string
var stdin bool
var outFile string
var getFromClipboard bool

// datatransformCmd represents the datatransform command
var datatransformCmd = &cobra.Command{
	Use:     "datatransform",
	Short:   "Data transformation utilities - JSON, CSV, TSV, Go Structs",
	Aliases: []string{"dt"},
	Long: `Some helpful datatransformation utilities for JSON, CSV, TSV, and Golang structs.

Current functionality includes:
* JSON > Go struct
* CSV > JSON
* CSV > Go struct`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("datatransform called")
	// },
}

// TODO - Make subcommands into input/output flags instead

func init() {
	cmdUtils.AddCommand(datatransformCmd)

	datatransformCmd.PersistentFlags().StringVar(&jsonString, "jsonString", "", "Json string to transform")
	datatransformCmd.PersistentFlags().StringVarP(&fileName, "fileName", "f", "", "Json string to transform")
	datatransformCmd.PersistentFlags().StringVarP(&webApi, "webApi", "a", "", "Json http API to get data and parse")
	datatransformCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "File to save output")
	datatransformCmd.PersistentFlags().BoolVarP(&stdin, "stdin", "i", false, "Used for piping in with stdin")
	datatransformCmd.PersistentFlags().BoolVarP(&getFromClipboard, "getFromClipboard", "c", false, "Get json from clipboard")
	// datatransformCmd.PersistentFlags().BoolVarP(&saveToClipboard, "saveToClipboard", "p", false, "Save output to clipboard")
	datatransformCmd.MarkFlagsMutuallyExclusive("jsonString", "fileName", "stdin", "getFromClipboard")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// datatransformCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// datatransformCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
