package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cliVersion = "0.3.0"

var (
	// The name of our config file, without the file extension because viper supports many different config file languages.
	defaultConfigFilename = "cldConf"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to STING_NUMBER.
	envPrefix = "CLD"

	// Replace hyphenated flag names with camelCase in the config file
	replaceHyphenWithCamelCase = true

	// outJSON flag
	OutJSON       bool
	DebugMode     bool
	ShowChangelog bool

	// // Only used when initially encrypting a previously unencrypted config file
	// InitialEncryptionOfUnencryptedConfigFile bool
)

//go:embed CHANGELOG.md
var ChangelogFile string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "cld",
	Version: cliVersion,
	Short:   "A brief description of your application",
	Long: `This CLI has been created to add additional functionality
to Azure CLI such as data
aggregation from multiple 'az' commands into a MongoDB Dababase, reporting,
and pulling data from both Azure DevOps and Azure,
as well as other functionality. AWS functionality is also being added.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well

		// if Changelog {
		// 	fmt.Println(ChangelogFile)
		// 	os.Exit(0)
		// }

		if DebugMode {
			fmt.Println("DEBUG MODE")

			cmd.DebugFlags()

			// Create a CPU profile file
			f, err := os.Create("profile.prof")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			// Start CPU profiling
			if err := pprof.StartCPUProfile(f); err != nil {
				panic(err)
			}
			defer pprof.StopCPUProfile()

			// Create a memory profile file
			memProfileFile, err := os.Create("mem.prof")
			if err != nil {
				panic(err)
			}
			defer memProfileFile.Close()

			// Write memory profile to file
			if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
				panic(err)
			}
			fmt.Println("Memory profile written to mem.prof")

			// Start tracing
			traceFile, err := os.Create("trace.out")
			if err != nil {
				panic(err)
			}
			defer traceFile.Close()

			if err := trace.Start(traceFile); err != nil {
				panic(err)
			}
			defer trace.Stop()
		}
		return InitializeConfig(cmd)
	},

	Run: func(cmd *cobra.Command, args []string) {
		if ShowChangelog {
			ViewChangelog(ChangelogFile)
		}

		if !ShowChangelog {
			cmd.Help()
		}
	},
}

func InitializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	// Set the base name of the config file, without the file extension.
	v.SetConfigName(defaultConfigFilename)

	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath("~/.config")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(envPrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&OutJSON, "outJSON", "j", false, "Output formatted to JSON")
	RootCmd.PersistentFlags().BoolVar(&DebugMode, "debug", false, "Debug mode creates trace logs for Golang pprof")
	RootCmd.Flags().BoolVar(&ShowChangelog, "changelog", false, "Shows Cloudini Changelog")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
