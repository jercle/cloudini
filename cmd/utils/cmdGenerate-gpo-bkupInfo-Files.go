package utils

import (
	"github.com/spf13/cobra"
)

var (
	rootPath         string
	printOne         bool
	gpoGuid          string
	gpoName          string
	backupId         string
	domain           string
	domainController string
)

// subsCmd represents the subs command
var cmdGenerateGpoBkupInfoFiles = &cobra.Command{
	Use:     "gpo-bkupinfo-files",
	Aliases: []string{"gpobuinf"},
	Short:   "Generate bkupInfo.xml files for AD Group Policy backups",
	Long: `
	Backup path must follow: ROOT/GPO_NAME/GPO_GUID/{BACKUP_ID}/

	--path/-p flag is for the root directory of all backups
	if -p is not provided, root will be the current directory

	You can also use the --print-one/-o flag to output one to stdout
	using the --gpoGuid, --gpoName, and --backupId flags.
	If gpoGuid and backupId are unknown, provide arbitrary uuids for template injection

Example:
  cld u gen gpobuinf -p ./gpo_updates`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(passwordLength)
		// pwd, err := lib.GenerateRandomString(passwordLength, passwordIncludeUpper, passwordIncludeNumbers, passwordIncludeSpecial)
		// lib.CheckFatalError(err)
		// fmt.Println(pwd)
		if printOne {
			GenerateGpoBkupInfoToStdOut(gpoName, gpoGuid, backupId, domain, domainController)
		} else {
			GenerateGpoBkupInfoFiles(rootPath)
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
	cmdGenerate.AddCommand(cmdGenerateGpoBkupInfoFiles)

	cmdGenerateGpoBkupInfoFiles.Flags().StringVarP(&rootPath, "path", "p", "./", "Root path of AD GPO backups")
	cmdGenerateGpoBkupInfoFiles.Flags().BoolVarP(&printOne, "print-one", "o", false, "Print one file to stdout, given the --gpoGuid, --gpoName, and --backupId flags")
	cmdGenerateGpoBkupInfoFiles.Flags().StringVar(&gpoGuid, "gpoGuid", "", "GPO GUID")
	cmdGenerateGpoBkupInfoFiles.Flags().StringVar(&gpoName, "gpoName", "", "GPO display name")
	cmdGenerateGpoBkupInfoFiles.Flags().StringVar(&backupId, "backupId", "", "GPO Backup ID GUID")
	cmdGenerateGpoBkupInfoFiles.Flags().StringVar(&domain, "domain", "", "GPO domain")
	cmdGenerateGpoBkupInfoFiles.Flags().StringVar(&domainController, "domainController", "", "GPO domain controller")
	cmdGenerateGpoBkupInfoFiles.MarkFlagsRequiredTogether("print-one", "gpoGuid", "gpoName", "backupId", "domain", "domainController")
}
