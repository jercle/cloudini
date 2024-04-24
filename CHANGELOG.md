# Change Log

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
