# Cloudini


This CLI has been created to add additional functionality to [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/) such as data aggregation from multiple `az` commands, reporting, and pulling data from both Azure DevOps and Azure.

Some of the reporting functionality is around container vulnerability scanning with the ability to install a web portal as an Azure Web App in development

This CLI is still very much under development, and functions with the `test-` or `dev-` prefixes are considered mid-development

Any help with testing would be greatly appreciated, as that area is my biggest weakness.

- [Cloudini](#cloudini)
  - [Getting started](#getting-started)
  - [Functionality](#functionality)
    - [Proxy settings (Windows only)](#proxy-settings-windows-only)


## Getting started
Copy [cldConf.json](https://raw.githubusercontent.com/jercle/cloudini/develop/cldConf.json) to `$HOME\.config\cld\` on mac/linux `%userprofile%\.config\cld\` on Windows

## Functionality
### Proxy settings (Windows only)
  * Set default proxy configured in cloudini
    * `cld.exe u win proxy -s`
  * Get currently configured proxy in Windows
    * `cld.exe u win proxy`
  * Open 'Internet Options' control panel cmdlet
    * `cld.exe u win proxy -o`
  * Remove proxy configuration
    * `cld.exe u win proxy -d`
  * Set proxy configuration from a proxyConfig object in the cloudini config file
    * `cld.exe u win proxy -s -n NAME`
