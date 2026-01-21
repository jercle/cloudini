package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/jercle/cloudini/lib"
)

func GeneratePassword(length int, includeUpper bool, includeNumbers bool, includeSpecial bool) string {
	pwd, err := lib.GenerateRandomString(length, includeUpper, includeNumbers, includeSpecial)
	lib.CheckFatalError(err)
	return pwd
}

//
//

func GenerateGpoBkupInfoFiles(gpoBackupPath string) {
	backups, err := os.ReadDir(gpoBackupPath)
	lib.CheckFatalError(err)

	for _, bu := range backups {
		if bu.Name() == "WMI_Filters" {
			continue
		}

		var backupInfo GroupPolicyObjectBackupInfo

		backupInfo.GPOName = bu.Name()
		cwd := filepath.Join(gpoBackupPath, backupInfo.GPOName)
		subDirs, err := os.ReadDir(cwd)
		lib.CheckFatalError(err)

		backupInfo.GPOGuid = subDirs[0].Name()
		gpoGuidDir := filepath.Join(gpoBackupPath, backupInfo.GPOName, backupInfo.GPOGuid)
		lib.CheckFatalError(err)

		gpoGuidSubDirs, err := os.ReadDir(gpoGuidDir)
		backupInfo.BackupId = gpoGuidSubDirs[0].Name()
		// fmt.Println("backupInfo.Name", backupInfo.GPOName)
		// fmt.Println("backupInfo.Guid", backupInfo.GPOGuid)
		// fmt.Println("backupInfo.BackupId", backupInfo.BackupId)

		tmplStr := `<BackupInst xmlns="http://www.microsoft.com/GroupPolicy/GPOOperations/Manifest">
	<GPOGuid><![CDATA[{ {{- .GPOGuid -}} }]]></GPOGuid>
	<GPODomain><![CDATA[ {{- .Domain -}} ]]></GPODomain>
	<GPODomainGuid><![CDATA[{  }]]></GPODomainGuid>
	<GPODomainController><![CDATA[ {{- .DomainController -}} ]]></GPODomainController>
	<BackupTime><![CDATA[2025-01-28T00:23:11]]></BackupTime>
	<ID><![CDATA[ {{- .BackupId -}} ]]></ID>
	<Comment><![CDATA[]]></Comment>
	<GPODisplayName><![CDATA[ {{- .GPOName -}} ]]></GPODisplayName>
</BackupInst>`

		tmpl, err := template.New("bkupInfo").Parse(tmplStr)
		lib.CheckFatalError(err)

		buInfoFilePath := filepath.Join(gpoBackupPath, backupInfo.GPOName, backupInfo.GPOGuid, backupInfo.BackupId) + "/bkupInfo.xml"
		fmt.Println("Creating file", buInfoFilePath)

		buInfoFile, err := os.Create(buInfoFilePath)
		lib.CheckFatalError(err)

		err = tmpl.Execute(buInfoFile, backupInfo)
		lib.CheckFatalError(err)

		buInfoFile.Close()
	}
}

//
//

func GenerateGpoBkupInfoToStdOut(gpoName string, gpoGuid string, backupId string, domain string, domainController string) {

	backupInfo := GroupPolicyObjectBackupInfo{
		GPOName:          gpoName,
		GPOGuid:          gpoGuid,
		BackupId:         backupId,
		Domain:           domain,
		DomainController: domainController,
	}

	tmplStr := `<BackupInst xmlns="http://www.microsoft.com/GroupPolicy/GPOOperations/Manifest">
  <GPOGuid><![CDATA[{ {{- .GPOGuid -}} }]]></GPOGuid>
  <GPODomain><![CDATA[ {{- .Domain -}} ]]></GPODomain>
  <GPODomainGuid><![CDATA[{  }]]></GPODomainGuid>
  <GPODomainController><![CDATA[ {{- .DomainController -}} ]]></GPODomainController>
  <BackupTime><![CDATA[2025-01-28T00:23:11]]></BackupTime>
  <ID><![CDATA[ {{- .BackupId -}} ]]></ID>
  <Comment><![CDATA[]]></Comment>
  <GPODisplayName><![CDATA[ {{- .GPOName -}} ]]></GPODisplayName>
</BackupInst>`

	tmpl, err := template.New("bkupInfo").Parse(tmplStr)
	lib.CheckFatalError(err)

	err = tmpl.Execute(os.Stdout, backupInfo)
	lib.CheckFatalError(err)

}

//
//

type GroupPolicyObjectBackupInfo struct {
	GPOName          string
	GPOGuid          string
	BackupId         string
	DomainController string
	Domain           string
}
