package utils

import (
	"fmt"
	"log"

	dt "github.com/jercle/cloudini/cmd/utils/datatransforms"
	"github.com/spf13/cobra"
)

// var filename string

// csv2structCmd represents the csv2struct command
var csv2structCmd = &cobra.Command{
	Use:     "csv2struct",
	Aliases: []string{"c2s"},
	Short:   "Converts CSV to a Golang Struct. Input selected using flags",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var opts dt.GenerateOptions
		jsonData := dt.CsvFileToJson(fileName)
		opts.JsonString = jsonData
		opts.JsonFile = ""
		opts.WebApi = ""
		opts.Stdin = stdin
		opts.GetFromClipboard = getFromClipboard

		code, err := dt.Generate(opts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(code)
	},
}

func init() {

	datatransformCmd.AddCommand(csv2structCmd)
	// csv2structCmd.Flags().StringVarP(&fileName, "filename", "f", "", "Csv file transform")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csv2structCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csv2structCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
