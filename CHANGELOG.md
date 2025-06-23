# Change Log

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
