package ado

// var pat = os.Getenv("DEVOPS_PAT")
// var org = os.Getenv("DEVOPS_ORG")
// var project = os.Getenv("DEVOPS_PROJECT")

func GetPipelines(org string, project string, pat string) []byte {
	urlString := AzureDevOpsBaseUrl + org + "/" + project + "/_apis/pipelines?" + AzureDevopsApiVersion
	data := azureDevOpsRestGet(urlString, pat)
	return data
}

// TODO: Get request for only failed pipelines

// Example response
// {
//     "count": 1,
//     "value": [
//         {
//             "_links": {
//                 "self": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3?revision=1"
//                 },
//                 "web": {
//                     "href": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_build/definition?definitionId=3"
//                 }
//             },
//             "url": "https://dev.azure.com/stkcat/e55895d8-4f47-4a25-b598-34e5f97dcd43/_apis/pipelines/3?revision=1",
//             "id": 3,
//             "revision": 1,
//             "name": "pipeline-name",
//             "folder": "\\"
//         }
//     ]
// }
