package ado

func getPipelineRuns(pipelineId string, pat string, project string, org string) []byte {
	urlString := AzureDevOpsBaseUrl + org + "/" + project + "/_apis/pipelines/" + pipelineId + "/runs?" + AzureDevopsApiVersion
	data := azureDevOpsRestGet(urlString, pat)
	return data
}



// Example response
// {
//     "count": 1,
//     "value": [
//         {
//             "_links": {
//                 "self": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3/runs/30"
//                 },
//                 "web": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_build/results?buildId=30"
//                 },
//                 "pipeline.web": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_build/definition?definitionId=3"
//                 },
//                 "pipeline": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3?revision=1"
//                 }
//             },
//             "templateParameters": {},
//             "pipeline": {
//                 "url": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3?revision=1",
//                 "id": 3,
//                 "revision": 1,
//                 "name": "pipeline-name",
//                 "folder": "\\"
//             },
//             "state": "completed",
//             "result": "succeeded",
//             "createdDate": "2023-09-28T06:23:26.5354401Z",
//             "finishedDate": "2023-09-28T06:24:00.8283057Z",
//             "url": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3/runs/30",
//             "id": 30,
//             "name": "TFP - keyvault"
//         }
//     ]
// }
