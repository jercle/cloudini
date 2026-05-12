package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	dt "github.com/jercle/cloudini/cmd/utils/datatransforms"
)

// var jsonString io.Reader
// adoCmd represents the ado command

// var saveToClipboard bool

// json2structCmd represents the json2struct command
var json2structCmd = &cobra.Command{
	Use:     "json2struct",
	Aliases: []string{"j2s"},
	Short:   "Converts JSON to a Golang Struct. Input selected using flags",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("ado called")
		// jsonString := `{"age":38,"user_height_m":1.7,"favoriteFoods":["cake"]}`
		// '{"age":38,"user_height_m":1.7,"favoriteFoods":["cake"]}'

		var opts dt.GenerateOptions

		opts.JsonString = jsonString
		opts.JsonFile = fileName
		opts.WebApi = webApi
		opts.Stdin = stdin
		opts.GetFromClipboard = getFromClipboard

		// fmt.Println(fileName)

		code, err := dt.Generate(opts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(code)

		if outFile != "" {
			os.WriteFile(outFile, []byte(code), os.ModePerm)
		}

		// if saveToClipboard {
		// 	clipboard.Write(clipboard.FmtText, []byte("testion"))
		// }
	},
}

func init() {
	datatransformCmd.AddCommand(json2structCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// json2structCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// json2structCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
