param (
    [string]$Token,
    [string]$Organization,
    [string]$AppId
)

#Install-Module -Name ExchangeOnlineManagement

Import-Module -Name ExchangeOnlineManagement



Connect-ExchangeOnline -ShowBanner:$false -Organization $Organization -AppId $AppId -AccessToken $Token

$inputData = [Console]::In.ReadToEnd()
$Mailboxes = $inputData | ConvertFrom-Json

$processedMailboxes = @()

foreach ($Mailbox in $Mailboxes) {
    $Mailbox | Add-Member -MemberType NoteProperty -Name "Permissions" -Value @()
    $Mailbox | Add-Member -MemberType NoteProperty -Name "MailboxPermissions" -Value @()
    $Mailbox | Add-Member -MemberType NoteProperty -Name "UserPermissions" -Value @()
    $MailboxPermissions = @(Get-MailboxPermission -Identity $Mailbox.UserPrincipalName -ResultSize Unlimited |
        Where-Object { $_.User -ne "NT AUTHORITY\SELF" } |
        Select-Object User, IsOwner, AccessRights)

    $UserPermissions = @(Get-RecipientPermission -Identity $Mailbox.UserPrincipalName -ResultSize Unlimited |
        Where-Object { $_.Trustee -notlike "NT AUTHORITY\SELF" } |
        Select-Object Trustee, AccessRights)

    $Mailbox.MailboxPermissions = $MailboxPermissions
    $Mailbox.UserPermissions = $UserPermissions

    foreach ($Perm in $MailboxPermissions) {
        $Mailbox.Permissions += [PSCustomObject]@{
            User       = $Perm.User
            Permission = "FullAccess"
            IsOwner    = $Perm.IsOwner
        }
    }
    foreach ($Perm in $UserPermissions) {
        $Mailbox.Permissions += [PSCustomObject]@{
            User       = $Perm.Trustee
            Permission = "SendAs"
        }
    }
    foreach ($Perm in $Mailbox.GrantSendOnBehalfTo) {
        $Mailbox.Permissions += [PSCustomObject]@{
            User       = $Perm
            Permission = "SendOnBehalf"
        }
    }

    $processedMailboxes += $Mailbox
}
Disconnect-ExchangeOnline -Confirm:$false
$processedMailboxes | Select-Object UserPrincipalName, Alias, DisplayName, EmailAddresses, RecipientTypeDetails, Permissions | ConvertTo-Json -Depth 4 -AsArray
# $processedMailboxes | ConvertTo-Json -Depth 4 -AsArray
