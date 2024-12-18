# Change Log

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
