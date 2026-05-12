Get-Mailbox -ResultSize Unlimited | ForEach-Object {
  [PSCustomObject]@{
     PrimarySmtpAddress = (Get-Mailbox $_.Identity).PrimarySmtpAddress.ToLower()
     MailboxTotalItemSize = (Get-MailboxStatistics $_.Identity).TotalItemSize
  }
} | Sort-Object TotalItemSize | Format-Table -AutoSize -HideTableHeaders
