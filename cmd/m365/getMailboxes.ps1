param (
    [string]$Token,
    [string]$Organization,
    [string]$AppId
)


#Install-Module -Name ExchangeOnlineManagement

Import-Module -Name ExchangeOnlineManagement

Connect-ExchangeOnline -ShowBanner:$false -Organization $Organization -AppId $AppId -AccessToken $Token

# Get-ExoMailbox -ResultSize 5 -RecipientTypeDetails SharedMailbox -Properties RecipientTypeDetails
Get-ExoMailbox -ResultSize Unlimited -RecipientTypeDetails SharedMailbox -Properties RecipientTypeDetails
| ConvertTo-Json
