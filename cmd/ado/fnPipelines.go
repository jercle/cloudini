package ado

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelines"
)

type ListProjectsResponse struct {
	Value             []Project
	ContinuationToken string
}

type Project struct {
	ID             string    `json:"id"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
	Name           string    `json:"name"`
	Revision       float64   `json:"revision"`
	State          string    `json:"state"`
	URL            string    `json:"url"`
	Visibility     string    `json:"visibility"`
}

type PipelineResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Web struct {
			Href string `json:"href"`
		} `json:"web"`
	} `json:"_links"`
	Folder   string  `json:"folder"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Revision float64 `json:"revision"`
	URL      string  `json:"url"`
}

type Pipeline struct {
	Name    string        `json:"name"`
	ID      int           `json:"id"`
	Project string        `json:"project"`
	Folder  string        `json:"folder"`
	WebUrl  string        `json:"webUrl"`
	Runs    []PipelineRun `json:"runs"`
}

type PipelineRunResponse struct {
	Links struct {
		Pipeline struct {
			Href string `json:"href"`
		} `json:"pipeline"`
		Pipeline_Web struct {
			Href string `json:"href"`
		} `json:"pipeline.web"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Web struct {
			Href string `json:"href"`
		} `json:"web"`
	} `json:"_links"`
	CreatedDate  time.Time `json:"createdDate"`
	FinishedDate time.Time `json:"finishedDate"`
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Pipeline     struct {
		Folder   string `json:"folder"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Revision int    `json:"revision"`
		WebUrl   string `json:"url"`
	} `json:"pipeline"`
	Result             string   `json:"result"`
	State              string   `json:"state"`
	TemplateParameters struct{} `json:"templateParameters"`
	WebUrl             string   `json:"url"`
}

type PipelineRun struct {
	Name         string    `json:"name"`
	ID           int       `json:"id"`
	CreatedDate  time.Time `json:"createdDate"`
	FinishedDate time.Time `json:"finishedDate"`
	Result       string    `json:"result"`
	State        string    `json:"state"`
	WebUrl       string    `json:"webUrl"`
	Pipeline     struct {
		Name    string `json:"name"`
		ID      int    `json:"id"`
		Project string `json:"project"`
		Folder  string `json:"folder"`
		WebUrl  string `json:"webUrl"`
	}
}

func (pipeline *Pipeline) GetRuns(ctx context.Context, connection *azuredevops.Connection) []PipelineRun {
	var (
		pipelineRuns         []PipelineRun
		pipelineRunsResponse []PipelineRunResponse
		options              pipelines.ListRunsArgs
	)

	pipelinesClient := pipelines.NewClient(ctx, connection)

	options.Project = &pipeline.Project
	options.PipelineId = &pipeline.ID

	runs, err := pipelinesClient.ListRuns(ctx, options)
	lib.CheckFatalError(err)

	jsonResponse, err := json.Marshal(runs)
	lib.CheckFatalError(err)

	err = json.Unmarshal(jsonResponse, &pipelineRunsResponse)
	lib.CheckFatalError(err)

	for _, run := range pipelineRunsResponse {
		var currentRun PipelineRun

		currentRun.Name = run.Name
		currentRun.ID = run.ID
		currentRun.CreatedDate = run.CreatedDate
		currentRun.FinishedDate = run.FinishedDate
		currentRun.Result = run.Result
		currentRun.State = run.State
		currentRun.WebUrl = run.Links.Web.Href
		currentRun.Pipeline.Name = pipeline.Name
		currentRun.Pipeline.ID = pipeline.ID
		currentRun.Pipeline.Project = pipeline.Project
		currentRun.Pipeline.Folder = pipeline.Folder
		currentRun.Pipeline.WebUrl = pipeline.WebUrl

		pipelineRuns = append(pipelineRuns, currentRun)
	}

	return pipelineRuns
}

func (project *Project) GetPipelines(ctx context.Context, connection *azuredevops.Connection) []Pipeline {
	var (
		allPipelines []Pipeline
		options      pipelines.ListPipelinesArgs
	)

	pipelinesClient := pipelines.NewClient(ctx, connection)

	options.Project = &project.Name

	pipelinesList, err := pipelinesClient.ListPipelines(ctx, options)
	lib.CheckFatalError(err)

	for _, pl := range *pipelinesList {
		var jsonPipeline PipelineResponse
		var thisPipeline Pipeline

		marshPl, err := json.Marshal(pl)
		lib.CheckFatalError(err)
		err = json.Unmarshal(marshPl, &jsonPipeline)
		lib.CheckFatalError(err)

		thisPipeline.Name = jsonPipeline.Name
		thisPipeline.WebUrl = jsonPipeline.Links.Web.Href
		thisPipeline.ID = jsonPipeline.ID
		thisPipeline.Folder = jsonPipeline.Folder
		thisPipeline.Project = project.Name

		allPipelines = append(allPipelines, thisPipeline)
	}

	return allPipelines
}

func getAllRuns(ctx context.Context, connection *azuredevops.Connection) []PipelineRun {
	var (
		allPipelines []Pipeline
		allRuns      []PipelineRun
		mutex        sync.Mutex
		wg           sync.WaitGroup
	)

	allProjects := getProjects(ctx, connection)

	for _, project := range allProjects {
		allPipelines = append(allPipelines, project.GetPipelines(ctx, connection)...)
	}

	for _, pl := range allPipelines {
		wg.Add(1)
		go func() {
			plRuns := pl.GetRuns(ctx, connection)
			mutex.Lock()
			allRuns = append(allRuns, plRuns...)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return allRuns
}

func getAllPipelinesWithRuns(ctx context.Context, connection *azuredevops.Connection) []Pipeline {
	var (
		allPipelines         []Pipeline
		allPipelinesWithRuns []Pipeline
		mutex                sync.Mutex
		wg                   sync.WaitGroup
	)

	allProjects := getProjects(ctx, connection)

	for _, project := range allProjects {
		allPipelines = append(allPipelines, project.GetPipelines(ctx, connection)...)
	}

	for _, pl := range allPipelines {
		// var thisPipeline Pipeline
		wg.Add(1)
		go func() {
			plRuns := pl.GetRuns(ctx, connection)
			pl.Runs = plRuns
			mutex.Lock()
			allPipelinesWithRuns = append(allPipelinesWithRuns, pl)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return allPipelinesWithRuns
}

func getProjects(ctx context.Context, connection *azuredevops.Connection) []Project {
	var allProjects []Project

	coreClient, err := core.NewClient(ctx, connection)
	lib.CheckFatalError(err)

	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	lib.CheckFatalError(err)

	for responseValue != nil {
		// Log the page of team project names
		for _, project := range (*responseValue).Value {
			var currentProject Project
			jsonData, err := json.Marshal(project)
			lib.CheckFatalError(err)

			err = json.Unmarshal(jsonData, &currentProject)
			lib.CheckFatalError(err)

			allProjects = append(allProjects, currentProject)
		}

		// if continuationToken has a value, then there is at least one more page of projects to get
		if responseValue.ContinuationToken != "" {

			continuationToken, err := strconv.Atoi(responseValue.ContinuationToken)
			if err != nil {
				log.Fatal(err)
			}

			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}
	return allProjects
}
