package utils

import (
	"fmt"

	"github.com/antonmedv/clipboard"
	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	passwordLength         int
	passwordIncludeUpper   bool
	passwordIncludeNumbers bool
	passwordIncludeSpecial bool
	copyToClipboard        bool
)

// subsCmd represents the subs command
var cmdGeneratePwd = &cobra.Command{
	Use:     "password",
	Aliases: []string{"pwd"},
	Short:   "Generate random secure password",
	Long: `Generates a random password string, with the default length of 20 characters.

To generate a password of a different length, provide the desired length with the '-p' flag.

If you want to include a character set, use the approprate flag.
You can include multiple sets without individual flags
Example:
'-un' (enable uppercase and numbers)`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(passwordLength)
		pwd, err := lib.GenerateRandomString(passwordLength, passwordIncludeUpper, passwordIncludeNumbers, passwordIncludeSpecial)
		lib.CheckFatalError(err)
		fmt.Println(pwd)

		if copyToClipboard {
			clipboard.WriteAll(pwd)
		}
	},
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {

	// },

	// PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {
	// 	// set resourceGroup flag as required for subcommands of this
	// 	azCmd.MarkPersistentFlagRequired("resourceGroup")
	// 	// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
	// 	return cmd.InitializeConfig(ccmd)
	// },
}

func init() {
	cmdGenerate.AddCommand(cmdGeneratePwd)

	cmdGeneratePwd.Flags().IntVarP(&passwordLength, "length", "l", 20, "Length of password to be generated")
	cmdGeneratePwd.Flags().BoolVarP(&passwordIncludeUpper, "include-uppercase", "u", false, "Include uppercase characters")
	cmdGeneratePwd.Flags().BoolVarP(&passwordIncludeNumbers, "include-numbers", "n", false, "Include uppercase characters")
	cmdGeneratePwd.Flags().BoolVarP(&passwordIncludeSpecial, "include-special", "s", false, "Include uppercase characters")
	cmdGeneratePwd.Flags().BoolVarP(&copyToClipboard, "copy-to-clipboard", "c", false, "Copies generated value to system clipboard")
}
