package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
)

func GetCertAuthCertInfoFromFile(path string) (processedItems []CertAuthorityCertInfo) {
	file, err := ReadFileUTF16(path)
	CheckFatalError(err)

	if len(file) == 0 {
		return nil
	}
	// fmt.Println(path)
	var fileData []map[string]string
	err = json.Unmarshal(file, &fileData)
	CheckFatalError(err)

	stat, _ := os.Stat(path)
	fileModTime := stat.ModTime()

	for _, item := range fileData {
		curr := make(map[string]interface{})
		for key, val := range item {

			if key == "Archived Key" ||
				key == "Key Recovery Agent Hashes" ||
				key == "Old Certificate" ||
				key == "Certificate Hash" ||
				key == "Template General Flags" ||
				key == "Template Private Key Flags" ||
				key == "Template Enrollment Flags" ||
				key == "Issued Subject Key Identifier" ||
				key == "Public Key Algorithm Parameters" ||
				key == "Binary Certificate" ||
				key == "Binary Public Key" ||
				key == "Issued Binary Name" ||
				key == "Request Binary Name" ||
				key == "Binary Request" ||
				val == "EMPTY" {
				continue
			}

			camelKey := strcase.ToLowerCamel(key)
			if camelKey == "certificateEffectiveDate" ||
				camelKey == "certificateExpirationDate" ||
				camelKey == "requestResolutionDate" ||
				camelKey == "requestSubmissionDate" {
				valDate, err := time.Parse("1/2/2006 3:04 PM", val)
				CheckFatalError(err)
				curr[camelKey] = valDate
			} else {
				changedVal := strings.ReplaceAll(val, "\\", "/")
				curr[camelKey] = changedVal
			}
			if curr["revocationDate"] == nil {
				curr["crlPartitionIndex"] = nil
			}
			if curr["publishExpiredCertificateInCrl"] == "0" {
				curr["publishExpiredCertificateInCrl"] = false
			} else if curr["publishExpiredCertificateInCrl"] == "1" {
				curr["publishExpiredCertificateInCrl"] = true
			}
		}

		currStr, _ := json.Marshal(curr)
		var currData CertAuthorityCertInfo
		err = json.Unmarshal(currStr, &currData)
		CheckFatalError(err)

		currData.LastServerSync = fileModTime

		fileName := filepath.Base(path)
		fnSplit := strings.Join(strings.Split(fileName, "-")[1:], "-")
		fnSplit = strings.Split(fnSplit, ".")[0]
		tNameAndHostname := strings.Split(fnSplit, "_")
		currData.TenantName, currData.CertificateAuthorityName = tNameAndHostname[0], tNameAndHostname[1]

		processedItems = append(processedItems, currData)
	}

	return
}

//
//

func GetServerCertInfoFromFile(path string) (processedItems []ServerCertInfo) {
	file, err := ReadFileUTF16(path)
	CheckFatalError(err)

	// fmt.Println(path)
	var fileData []map[string]interface{}
	err = json.Unmarshal(file, &fileData)
	CheckFatalError(err)

	stat, _ := os.Stat(path)
	fileModTime := stat.ModTime()

	for _, item := range fileData {
		curr := make(map[string]interface{})

		for key, val := range item {
			if key == "PSProvider" ||
				key == "PSDrive" ||
				key == "PSIsContainer" ||
				// key == "PSParentPath" ||
				// key == "PSPath" ||
				key == "PSChildName" {
				continue
			}

			if key == "NotBefore" || key == "NotAfter" {
				valStr := strings.TrimPrefix(val.(string), "/Date(")
				valStr = strings.TrimSuffix(valStr, ")/")
				valInt, err := strconv.ParseInt(valStr, 10, 64)
				CheckFatalError(err)
				valDate := time.UnixMilli(valInt)
				curr[key] = valDate
			} else if key == "PSParentPath" {
				parentPath := strings.Split(val.(string), "::")[1]
				parentPathSpl := strings.Split(parentPath, "\\")
				curr["parentPath"] = strings.Join(parentPathSpl, "/")
			} else {
				curr[key] = val
			}
		}

		var currData ServerCertInfo
		jsonStr, _ := json.Marshal(curr)
		err = json.Unmarshal(jsonStr, &currData)
		CheckFatalError(err)

		fileName := filepath.Base(path)
		fnSplit := strings.Join(strings.Split(fileName, "-")[1:], "-")
		fnSplit = strings.Split(fnSplit, ".")[0]
		tNameAndHostname := strings.Split(fnSplit, "_")

		hostnameLower := strings.ToLower(tNameAndHostname[1])

		currData.TenantName = &tNameAndHostname[0]
		currData.TenantNames = append(currData.TenantNames, tNameAndHostname[0])
		currData.PulledFromServer = &hostnameLower
		currData.LastServerSync = fileModTime

		if len(*currData.EnhancedKeyUsageList) == 0 {
			*currData.EnhancedKeyUsageList = nil
		}

		if currData.EnrollmentPolicyEndPoint.AuthenticationType == 0 && currData.EnrollmentPolicyEndPoint.URL == nil {
			currData.EnrollmentPolicyEndPoint = nil
		}

		if currData.EnrollmentServerEndPoint.AuthenticationType == 0 && currData.EnrollmentServerEndPoint.URL == nil {
			currData.EnrollmentServerEndPoint = nil
		}

		currData.RawData = nil
		currData.IssuerName.RawData = nil
		currData.SubjectName.RawData = nil

		currDataStr, _ := json.Marshal(currData)
		var processedItem ServerCertInfo
		json.Unmarshal(currDataStr, &processedItem)
		processedItems = append(processedItems, processedItem)
	}

	return
}

//
//

func GetCertInfoFromFiles(basePath string, outputPath string) (caCertInfo []CertAuthorityCertInfo, serverCertInfo []ServerCertInfo) {
	paths := GetFullFilePaths(basePath)

	for _, path := range paths {

		if strings.Contains(path, "caCertList") {
			curr := GetCertAuthCertInfoFromFile(path)
			caCertInfo = append(caCertInfo, curr...)
		} else if strings.Contains(path, "serverCertList") {
			curr := GetServerCertInfoFromFile(path)
			serverCertInfo = append(serverCertInfo, curr...)
		}
	}

	if outputPath != "" {
		if _, err := os.Stat(outputPath); err != nil {
			os.MkdirAll(outputPath, os.ModePerm)
		}
		caCertInfoStr, _ := json.MarshalIndent(caCertInfo, "", "  ")
		serverCertInfoStr, _ := json.MarshalIndent(serverCertInfo, "", "  ")
		os.WriteFile(outputPath+"/caCertInfo.json", caCertInfoStr, 0644)
		os.WriteFile(outputPath+"/serverCertInfo.json", serverCertInfoStr, 0644)
	}

	return
}

//
//

func RelateCertAuthCertsToServerCerts(caCertInfo []CertAuthorityCertInfo, serverCertInfo []ServerCertInfo) (caCertInfoWithRelations []CertAuthorityCertInfo, serverCertInfoWithRelations []ServerCertInfo) {
	caCertsBySerialNumber := make(map[string]CertAuthorityCertInfo)
	serverCertsBySerialNumber := make(map[string]ServerCertInfo)

	for i, caci := range caCertInfo {
		if _, ok := caCertsBySerialNumber[caci.SerialNumber]; ok {
			fmt.Println("whoopsie! caCertsBySerialNumber")
			fmt.Println(i)
			os.Exit(1)
		}
		curr := caci
		curr.SerialNumber = strings.ToLower(caci.SerialNumber)
		caCertsBySerialNumber[curr.SerialNumber] = curr
	}

	for _, sci := range serverCertInfo {

		if sci.SerialNumber == "01" || sci.SerialNumber == "00" {
			continue
		}
		curr := sci
		curr.SerialNumber = strings.ToLower(sci.SerialNumber)
		curr.Thumbprint = strings.ToLower(sci.Thumbprint)

		// pfs := *sci.PulledFromServer + ":" + *sci.ParentPath

		pfs := ServerCertInfoServersPulledFrom{
			ServerName:       *sci.PulledFromServer,
			CertificatePaths: []string{*sci.ParentPath},
		}

		// var test bool
		// var foundData ServerCertInfo
		if data, ok := serverCertsBySerialNumber[curr.SerialNumber]; ok {
			// fmt.Println(data)
			// foundData = data
			// if curr.SerialNumber == "31360b71167360af4f00f3e3e4d0c213" {
			// 	test = true
			// }
			// fmt.Println("ok")
			curr.ServersPulledFrom = data.ServersPulledFrom
			curr.TenantNames = data.TenantNames
			// }
		}

		// fmt.Println(curr.TenantNames)
		// os.Exit(0)

		matchedSpf := false
		for index, spf := range curr.ServersPulledFrom {
			currSpf := spf
			if spf.ServerName == pfs.ServerName {
				matchedSpf = true
				currSpf.CertificatePaths = append(currSpf.CertificatePaths, pfs.CertificatePaths...)
			}
			curr.ServersPulledFrom[index] = currSpf
		}
		if !matchedSpf {
			curr.ServersPulledFrom = append(curr.ServersPulledFrom, pfs)
		}

		// for _, tn := range curr.TenantNames {
		// 	// currTn := tn
		// 	// if tn != sci.TenantName {
		// 	// 	matchedTNs = true
		// 	// 	currSpf.CertificatePaths = append(currSpf.CertificatePaths, pfs.CertificatePaths...)
		// 	// }

		if !slices.Contains(curr.TenantNames, *curr.TenantName) {
			curr.TenantNames = append(curr.TenantNames, *curr.TenantName)
		}
		// curr.TenantNames[index] = currSpf
		// }
		// if !matchedTNs {
		// 	curr.TenantNames = append(curr.TenantNames, tn)
		// }

		curr.PulledFromServer = nil
		curr.TenantName = nil
		curr.ParentPath = nil

		if relatedCertAuth, ok := caCertsBySerialNumber[curr.SerialNumber]; ok {
			curr.RelatedCertAuthData = &relatedCertAuth
		}
		curr.ID = curr.SerialNumber
		serverCertsBySerialNumber[curr.SerialNumber] = curr
		// serverCertInfoWithRelations = append(serverCertInfoWithRelations, curr)

		// if test && curr.SerialNumber == "31360b71167360af4f00f3e3e4d0c213" {
		// JsonMarshalAndPrint(foundData)
		// JsonMarshalAndPrint(sci)
		// JsonMarshalAndPrint(curr)
		// os.Exit(0)
		// }
	}

	// JsonMarshalAndPrint(serverCertsBySerialNumber["31360b71167360af4f00f3e3e4d0c213"])

	for _, cert := range serverCertsBySerialNumber {
		serverCertInfoWithRelations = append(serverCertInfoWithRelations, cert)
	}

	// serverCertInfoWithRelations = append(serverCertInfoWithRelations, curr)

	for _, caci := range caCertsBySerialNumber {
		curr := caci
		curr.ID = curr.SerialNumber
		curr.RelatedServersCertUsedOn = serverCertsBySerialNumber[caci.SerialNumber].ServersPulledFrom
		caCertInfoWithRelations = append(caCertInfoWithRelations, curr)
	}

	return
}
