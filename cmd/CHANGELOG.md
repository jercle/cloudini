# Change Log

## [0.3.44]
* Update deps

## [0.3.43]
* Skip B2C tenants for IP address fetching
* Add website certificate checking functionality
* Update "TenantName" field on server certs to cater for identical certificates across multiple tenants

## [0.3.42]
* Add azure/network/P2SVpn

## [0.3.41]
* Bugfix: Incorrect order of string/prefix when checking mongodb string

## [0.3.40]
* Update to Golang 1.25.5
* Update deps

## [0.3.39]
* Add handling of mongodb with srv connection string

## [0.3.38]
* Remove UnknownFields from azure.AzureAlertProcessed

## [0.3.37]
* Add filter to Support Alerts to skip closed alerts

## [0.3.36]
* Bugfix: Update GetAzureWorkbookAlerts to handle LinkToFilteredSearchResultsAPI not existing

## [0.3.35]
* Embed timezone data

## [0.3.34]
* Bugfix: Handling timezone in go alpine container

## [0.3.33]
* Bugfix: Update timezone to Australia/Sydney

## [0.3.32]
* Bugfix: Change ordering of InitConfig dir checks

## [0.3.31]
* Update handling Log Analytics queries to parse as local time if 'datetime_utc_to_local' is present in the query

## [0.3.30]
* Update lib.wrapperFuncs to use jsonv2
* Add functionality to get all resource, resource group, subscription, and management group role assignments for tenant

## [0.3.29]
* Bugfix: Fix Log Analytics query

## [0.3.28]
* Update deps
* Update fetching Resource SKUs
* B2C user fetching now includes extensions with prefix of 'extension' and suffixes of 'lastLogonTime' and 'passwordResetOn'

## [0.3.27]
* Update Managed Identity Env Var to AZURE_APPCONFIG_MANAGED_IDENTITY

## [0.3.26]
* Enable the use of Managed Identities for App Config using AZURE_APPCONFIG_USE_MANAGED_IDENTITY

## [0.3.25]
* Update dependencies

## [0.3.24]
* Bugfix: Fix env var name for AZURE_APPCONFIG_ENDPOINT

## [0.3.23]
* Add functionality to provide an Azure App Config label env var, which will then merge values with the label over the initial fetch if the env var exists.

## [0.3.22]
* Add using Azure App Configuration for config

## [0.3.20]
* Update Azure Log Analytics API endpoint to api.loganalytics.azure.com
* Update deps

## [0.3.19]
* Add getting and upserting IP Address Blocks to see blocks of IP Addresses used by vnets for given CIDR Blocks
* Update deps

## [0.3.18]
* Update ConvertLAResultToAWSIngestCounts to handle null values
* Update deps

## [0.3.17]
* Update subnet calc
* Update handling vnets and subnets when fetching resources, adding extra IP and CIDR related fields
* Update deps

## [0.3.16]
* Breakfix: remove duplicated flag in windows/cmdCert templated file
* Update deps

## [0.3.15]
* Handle empty curr.LinkToFilteredSearchResultsAPI string in azure.GetAzureWorkbookAlerts
* Update deps

## [0.3.14]
* Change handling of subnets when fetching all resources, breaking them out to their own resources
* Add CertificateManagement config field and update UpdateAllCertInfo to use

## [0.3.13]
* Bugfix: Update Log Analytics query token type

## [0.3.12]
* Bugfix: Change queryError type to string

## [0.3.11]
* Handle errors with Log Analytics alert queries

## [0.3.10]
* Bugfix: Change Tags prop on all Azure types to not omitempty

## [0.3.9]
* Bugfix: Fix public ip address query

## [0.3.8]
* Bugfix: Resources initially fetched from Azure do not have ExistsInAzure set to true

## [0.3.7]
* Bugfix: UpdateAzureResourceRelations failing on sources property
* Bugfix: UpdateResourcesNotExistInAzure not correctly applying ExistsInAzure prop
* Add back updateAzureResourceRelations flag

## [0.3.6]
* Bugfix: updateResources not actually getting current resources
* Temporarily remove updateAzureResourceRelations flag

## [0.3.5]
* AzureResourceDetails.ExistsInAzure omitempty removed from json/bson
* Add AWSIngestCounts.LastDBSync

## [0.3.4]
* Bugfix: GetDataFromPackerLogfile - nil dereference in field selection. Added var declaration back.

## [0.3.3]
* Add aws log ingest count query config
* Add azure.RunLogAnalyticsQuery
* Add mongodb.UpdateAWSMonitoringData

## [0.3.2]
* Bugfix: AzureResourceProperties.Other field saving to MongoDB as binary data
* Add updateResources flag to mongodb.Update command
* Add unsetResourceField flag to mongodb.Clean command
* Update deps

## [0.3.1]
* Bugfix: Some resources showing in IP Address list when resource has no ip addresses

## [0.3.0]
* Fixes and major overhaul to IP address retrieval to get from all resources.

## [0.2.1]
* Add updateAll flag to mongo update command

## [0.2.0]
* Switched to using Homebrew Casks instead of formulae

## [0.1.56]
* Bugfix: Remove os.Exit and print statement from GetDataFromPackerLogfile which was used for testing json/v2

## [0.1.55]
* Begin updating funcs to use json/v2 package
* Update deps

## [0.1.54]
* Bugfix: Update GetAzureWorkbookAlerts with correct date parsing

## [0.1.53]
* Bugfix: GetAllB2CTenantUsers not handling paged results

## [0.1.52]
* Update command descriptions
* Add charmbracelet/fang wrapper
* Add Changelog viewer
* Colourize JSON ouput when showing config
* Add functionality to get alerts from Azure Resource Graph, added config option for query
* Update deps

## [0.1.51]
* Add UpdateB2CUsers function to mongo update

## [0.1.50]
* Fix: Update bson flag for ServerCertInfoServersPulledFrom type CertificatePaths property

## [0.1.49]
* Fix: Update how serversPulledFrom is stored for certs to now transform them more efficiently and update the ServerCertInfoServersPulledFrom type

## [0.1.48]
* Add "checkExchange" config option, update to GetMailboxStorageUsedAllConfiguredTenants to check for checkExchange option

## [0.1.47]
* Fix: Update how serversPulledFrom is stored for certs to now transform them correctly

## [0.1.46]
* Feature: Add mailbox statistics cmd and mongodb upsert of mb stats data

## [0.1.45]
* Bug: Fix certificate aggregation to correctly include all servers pulled from

## [0.1.44]
* Update type field AzureResourceProperties.Visibility to type 'any'
* Update deps

## [0.1.43]
* Update GetServicePrincipalToken to allow using certificate for App Reg authentication
* Update deps

## [0.1.42]
* Bugfix: Fix to handling lib.AzureMultiAuthTokenRequestOptions.NoCache

## [0.1.41]
* Add AD funcs

## [0.1.40]
* Bugfix: Fix UpdateImageDataWithBuildHostLogs to only include logs of images currently existing in Azure

## [0.1.39]
* Update UpdateAllGalleryImagesAndUpdateWithUsedByCitrix to clear collection first.
* Update deps

## [0.1.38]
* Add lib.SliceOfStringsToUnique
* Add lib.PrintSliceIntsWithIndexes
* Add lib.RemoveSubdirectoriesOfPath
* Skip caching new token if NoCache is configured.

## [0.1.37]
* Bugfix: Update filename concatenation in azure.UploadBlobFromString and azure.UploadBlobFromFile

## [0.1.36]
* Bugfix: typo in StorageBlobHttpGet

## [0.1.35]
* Update azure.StorageBlobHttpGet to use azure.StorageAccountUploadBlobOptions for options input

## [0.1.34]
* Add TenantId to app reg data
* Rename azure.UploadBlob to azure.UploadBlobFromFile
* Create azure.UploadBlobFromString to allow handling of in memory data uploads
* Update to Go 1.24.3
* Update deps

## [0.1.33]
* Bugfix: Add back type and subscription Name to IP address fetch
* Bugfix: Cost data when some subscription names are caps has been resolved.

## [0.1.32]
* Change cert upserts to clear collection before upserting current certificates
* Enhancement: Vastly improved upsert perf with CosmosDB by using unordered writes, and chunking requests
* Bugfix: Fix issue with computegalleryimage cmd to re-enable persistent flags at az cmd level

## [0.1.31]
* Feature: Add subName and change min tls prop name to storage acct min tls data

## [0.1.30]
* Feature: Add Mongo update to include storage acct min TLS versions

## [0.1.29]
* Feature: Add cmd to check Min TLS versions of storage accounts in all configured tenants

## [0.1.28]
* Bugfix: Update to azure.GetServicePrincipalToken to differentiate between read/write tokens when performing cache operations

## [0.1.27]
* Bugfix: Change subnetcalc to use "github.com/rmasci/ipsubnet" instead of "github.com/brotherpowers/ipsubnet" for the time being, due to issues building 32-bit builds.

## [0.1.26]
* Add function for updating cert info
* Begin improving how we get all IP addresses
* Add "NoCache" option to getting tokens
* Update deps

## [0.1.25]
* Bugfix: fixed issue using mongo bulkwrite and updateone in same func
* Update deps

## [0.1.24]
* Add xmltags utility func and update deps

## [0.1.23]
* Fix bug: update bson flags so $connections is now _connections

## [0.1.22]
* Fix bug with function for getting IP addresses from Azure

## [0.1.21]
* Update deps, update Go to 1.24
* Add util/networking and subnet calc cmd
* Add mongodb/UpdateIpAddresses command
* Update bson flags so $schema is now _schema for CosmosDB compatibility

## [0.1.20]
* Remove tenantName check in cmdComputeGalleryImage

## [0.1.19]
* Add subscriptionId and resourceGroup flags to cmdComputeGalleryImage

## [0.1.18]
* Update deps
* Add tenantId back as persistant flag to azCmd

## [0.1.17]
* Update deps
* Add clientId and clientSecret flags back as persistant flags to azCmd

## [0.1.16]
* Update deps

## [0.1.15]
* Update deps
* Fix up bkupInfo generator cmd
* Add packer funcs, vm control commands, update auth caching, add packer build logs to image data in db, bkupInfo.xml generation for AD

## [0.1.14]
* Add commands for Azure B2C
* Add commands for Azure Entra
* Add commands for Azure Resources
* Update token caching
* Update functions for Azure Compute Gallery
* Add functions for Azure Storage Blobs
* Add functions for retrieving Azure Subscriptions
* Added HttpPatch functions for Azure
* Added Citrix command and functions for Machine Catalogs, Delivery Groups, User Info, http utils, and Auth
* Updated Cloudini Config map
* Added Mongodb commands and functions
* Added initial commands for Certificate management
* Updated password generation
* Added command to clear cached tokens

## [0.1.13]
* Fix bug in retrieving latest ACG image with Azure API chunking responses after first 55 versons.

## [0.1.12]
* Add back checkVersionExists flag for checking image gallery versions after accidental removal in 0.1.11

## [0.1.11]
* Updates to azure compute gallery, azure auth, and entra funcs
* Add m365 funcs
* Bug fix for config file location checking

## [0.1.10]
* Add functions to check if image version exists
* Add vm image version list functionality
* Removal of stale code

## [0.1.9]
* Update latest acp image getter
* Bump deps

## [0.1.8]
*  azure compute gallery: Remove error when no attempting to increment versions where non exist

## [0.1.7]
*  azure compute gallery: Remove failure when no image versions

## [0.1.6]
*  Bump dep versions
*  Update azure.HttpGet to error on 404
*  Add GetImage function

## [0.1.5]
*  Bump dep versions
*  Add functions for Shared Image Gallery/Compute Gallery
*  Add incrementing Compute Gallery image patch version command

## [0.1.4]

## [0.1.3]
 * Bug fix: #37 - remove GetActiveSub() call in cmdAz.go init() - fixes #37
 * Update proxy functions
 * Refactoring

## [0.1.2]
 * https://github.com/jercle/cloudini/pull/1
 * Added submodules for generated Azure fake data
 * Added dev/ directory for testing feature dev

## [0.1.1]

* Bump dep versions
* Utils
  * Feature: Updates and refactors to windows registry functions for getting installed apps
  * Feature: Add windows registry functions for updating proxy settings
* Azure
  * Feature: Log analytics functions
  * Feature: NSG flow log IP search
    * Bug: Using bad method of file iteration, greatly increased performance
  * Feature: Authentication improvements for multiple tenants
  * Feature: Standardised Azure REST API requests
  * Feature: Begin working on Azure Container Registry cleanup functionality
    * Will list all images, and prune all no longer in use with parameters
  * Feature: Retrieve Entra roles
* Azure DevOps
  * Feature: Add pipeline search
  * Feature: Begin working on listing top run pipelines
* Config
  * Feature: Added config functionality, intialisation, and structs
* Maintenance: Refactoring
