package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

		currData.TenantName = tNameAndHostname[0]
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
		if sci.SerialNumber == "01" {
			continue
		}
		curr := sci
		curr.SerialNumber = strings.ToLower(sci.SerialNumber)
		curr.Thumbprint = strings.ToLower(sci.Thumbprint)
		curr.PulledFromServer = nil
		// spf := *sci.PulledFromServer + ":" + *sci.ParentPath

		spf := ServerCertInfoServersPulledFrom{
			ServerName:      *sci.PulledFromServer,
			CertificatePath: *sci.ParentPath,
		}

		if data, ok := serverCertsBySerialNumber[sci.SerialNumber]; ok {
			// fmt.Println(data)
			// os.Exit(0)
			// curr = data
			curr.ServersPulledFrom = append(data.ServersPulledFrom, spf)
		} else {
			curr.PulledFromServer = nil
			curr.ParentPath = nil
			curr.ServersPulledFrom = append(curr.ServersPulledFrom, spf)
			serverCertsBySerialNumber[sci.SerialNumber] = curr
		}

		if relatedCertAuth, ok := caCertsBySerialNumber[curr.SerialNumber]; ok {
			curr.RelatedCertAuthData = &relatedCertAuth
		}
		serverCertsBySerialNumber[curr.SerialNumber] = curr
		curr.ID = curr.SerialNumber
		serverCertInfoWithRelations = append(serverCertInfoWithRelations, curr)
	}

	for _, caci := range caCertsBySerialNumber {
		curr := caci
		curr.ID = curr.SerialNumber
		curr.RelatedServersCertUsedOn = serverCertsBySerialNumber[caci.SerialNumber].ServersPulledFrom
		caCertInfoWithRelations = append(caCertInfoWithRelations, curr)
	}

	return
}
