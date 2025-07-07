package ado

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/spf13/cobra"
)

var (
	listAllPipelines         bool
	listPipelineRuns         string
	listAllPipelinesWithRuns bool
)

// pipelinesCmd represents the pipelines command
var pipelinesCmd = &cobra.Command{
	Use:     "pipelines",
	Aliases: []string{"pl"},
	Short:   "Commands related to ADO pipelines",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("pipelines called")
		orgUrl := "https://dev.azure.com/" + devOpsOrg
		ctx := context.Background()
		var allPipelines []Pipeline
		connection := azuredevops.NewPatConnection(orgUrl, personalAccessToken)

		if listAllPipelines {
			allProjects := getProjects(ctx, connection)
			for _, project := range allProjects {
				allPipelines = append(allPipelines, project.GetPipelines(ctx, connection)...)
			}
			jsonData, err := json.MarshalIndent(allPipelines, "", "  ")
			lib.CheckFatalError(err)
			fmt.Println(string(jsonData))
		}

		if listAllPipelinesWithRuns {
			allPipelines = getAllPipelinesWithRuns(ctx, connection)
			jsonData, err := json.MarshalIndent(allPipelines, "", "  ")
			lib.CheckFatalError(err)
			fmt.Println(string(jsonData))
		}

		if listPipelineRuns != "" {
			fmt.Println("TODO: Implement list runs for specific pipeline")
		}
	},
}

func init() {
	adoCmd.AddCommand(pipelinesCmd)

	pipelinesCmd.Flags().BoolVarP(&listAllPipelines, "listAllPipelines", "l", false, "List all pipelines")
	pipelinesCmd.Flags().StringVarP(&listPipelineRuns, "listRuns", "r", "", "List all runs for a specific pipeline")
	pipelinesCmd.Flags().BoolVarP(&listAllPipelinesWithRuns, "listAllPipelinesWithRuns", "a", false, "List all pipelines including all runs for each pipeline")
	pipelinesCmd.MarkFlagsMutuallyExclusive("listAllPipelines", "listRuns", "listAllPipelinesWithRuns")
	pipelinesCmd.MarkFlagsOneRequired("listAllPipelines", "listRuns", "listAllPipelinesWithRuns")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelinesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelinesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO list running pipelines
// TODO check specific pipeline
