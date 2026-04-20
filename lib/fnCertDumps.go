package lib

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
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

// func GetServerCertInfoFromFileNew(path string) (processedItems []FormattedServerCertInfo) {
// 	file, err := ReadFileUTF16(path)
// 	CheckFatalError(err)

// 	// fmt.Println(path)
// 	var fileData []map[string]interface{}
// 	err = json.Unmarshal(file, &fileData)
// 	CheckFatalError(err)

// 	stat, _ := os.Stat(path)
// 	fileModTime := stat.ModTime()

// 	for _, item := range fileData {
// 		curr := make(map[string]interface{})

// 		for key, val := range item {
// 			if key == "PSProvider" ||
// 				key == "PSDrive" ||
// 				key == "PSIsContainer" ||
// 				// key == "PSParentPath" ||
// 				// key == "PSPath" ||
// 				key == "PSChildName" {
// 				continue
// 			}

// 			if key == "NotBefore" || key == "NotAfter" {
// 				valStr := strings.TrimPrefix(val.(string), "/Date(")
// 				valStr = strings.TrimSuffix(valStr, ")/")
// 				valInt, err := strconv.ParseInt(valStr, 10, 64)
// 				CheckFatalError(err)
// 				valDate := time.UnixMilli(valInt)
// 				curr[key] = valDate
// 			} else if key == "PSParentPath" {
// 				parentPath := strings.Split(val.(string), "::")[1]
// 				parentPathSpl := strings.Split(parentPath, "\\")
// 				curr["parentPath"] = strings.Join(parentPathSpl, "/")
// 			} else {
// 				curr[key] = val
// 			}
// 		}

// 		// jsonStrRaw, _ := json.Marshal(curr)
// 		// os.WriteFile("main-cert-raw.json", jsonStrRaw, 0644)

// 		var currData ServerCertInfo
// 		jsonStr, _ := json.Marshal(curr)
// 		err = json.Unmarshal(jsonStr, &currData)
// 		CheckFatalError(err)

// 		fileName := filepath.Base(path)
// 		fnSplit := strings.Join(strings.Split(fileName, "-")[1:], "-")
// 		fnSplit = strings.Split(fnSplit, ".")[0]
// 		tNameAndHostname := strings.Split(fnSplit, "_")

// 		hostnameLower := strings.ToLower(tNameAndHostname[1])

// 		currData.TenantName = &tNameAndHostname[0]
// 		currData.TenantNames = append(currData.TenantNames, tNameAndHostname[0])
// 		currData.PulledFromServer = &hostnameLower
// 		currData.LastServerSync = fileModTime

// 		// if len(*currData.EnhancedKeyUsageList) == 0 {
// 		// 	*currData.EnhancedKeyUsageList = nil
// 		// }

// 		// if currData.EnrollmentPolicyEndPoint.AuthenticationType == 0 && currData.EnrollmentPolicyEndPoint.URL == nil {
// 		// 	currData.EnrollmentPolicyEndPoint = nil
// 		// }

// 		// if currData.EnrollmentServerEndPoint.AuthenticationType == 0 && currData.EnrollmentServerEndPoint.URL == nil {
// 		// 	currData.EnrollmentServerEndPoint = nil
// 		// }

// 		// // currData.RawData = nil
// 		// currData.IssuerName.RawData = nil
// 		// currData.SubjectName.RawData = nil

// 		currDataStr, _ := json.Marshal(currData)
// 		var processedItem ServerCertInfo
// 		json.Unmarshal(currDataStr, &processedItem)
// 		processedItems = append(processedItems, processedItem)
// 	}

//		return
//	}
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

		// jsonStrRaw, _ := json.Marshal(curr)
		// os.WriteFile("main-cert-raw.json", jsonStrRaw, 0644)

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

		// currData.RawData = nil
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

func RelateCertAuthCertsToServerCertsNew(caCertInfo []CertAuthorityCertInfo, serverCertInfo []FormattedServerCertInfo) (caCertInfoWithRelations []CertAuthorityCertInfo, serverCertInfoWithRelations []FormattedServerCertInfo) {
	caCertsBySerialNumber := make(map[string]CertAuthorityCertInfo)
	serverCertsBySerialNumber := make(map[string]FormattedServerCertInfo)

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

		if sci.Serial == "01" || sci.Serial == "00" {
			continue
		}
		curr := sci
		curr.Serial = strings.ToLower(sci.Serial)
		curr.Thumbprint = strings.ToLower(sci.Thumbprint)

		// pfs := *sci.PulledFromServer + ":" + *sci.ParentPath

		pfs := ServerCertInfoServersPulledFrom{
			ServerName:       *sci.PulledFromServer,
			CertificatePaths: []string{*sci.ParentPath},
		}

		// var test bool
		// var foundData ServerCertInfo
		if data, ok := serverCertsBySerialNumber[curr.Serial]; ok {
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

		if relatedCertAuth, ok := caCertsBySerialNumber[curr.Serial]; ok {
			curr.RelatedCertAuthData = &relatedCertAuth
		}
		curr.ID = curr.Serial
		serverCertsBySerialNumber[curr.Serial] = curr
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

//
//

func decodeAlgorithm(code x509.PublicKeyAlgorithm) string {
	var s string
	switch code {
	case x509.RSA:
		s = "RSA"
	case x509.DSA:
		s = "DSA"
	case x509.ECDSA:
		s = "ECDSA"
	default:
		s = "oops"
	}
	return s
}

//
//

func GetCertExtensionFromOID(oid string) string {
	mapping := map[string]string{
		"1.3.6.1.4.1.311.20.2":  "szOID_ENROLL_CERTTYPE",
		"1.3.6.1.4.1.311.21.1":  "MS Certificate Services CA Version",
		"1.3.6.1.4.1.311.21.2":  "szOID_CERTSRV_PREVIOUS_CERT_HASH",
		"1.3.6.1.4.1.311.21.3":  "szOID_CRL_VIRTUAL_BASE",
		"1.3.6.1.4.1.311.21.4":  "szOID_CRL_NEXT_PUBLISH",
		"1.3.6.1.4.1.311.21.5":  "szOID_KP_CA_EXCHANGE",
		"1.3.6.1.4.1.311.21.6":  "szOID_KP_KEY_RECOVERY_AGENT",
		"1.3.6.1.4.1.311.21.7":  "szOID_CERTIFICATE_TEMPLATE",
		"1.3.6.1.4.1.311.21.8":  "szOID_ENTERPRISE_OID_ROOT",
		"1.3.6.1.4.1.311.21.9":  "szOID_RDN_DUMMY_SIGNER",
		"1.3.6.1.4.1.311.21.10": "szOID_APPLICATION_CERT_POLICIES",
		"1.3.6.1.4.1.311.21.11": "szOID_APPLICATION_POLICY_MAPPINGS",
		"1.3.6.1.4.1.311.21.12": "szOID_APPLICATION_POLICY_CONSTRAINTS",
		"1.3.6.1.4.1.311.21.13": "Attribute added to an certificate request when key archival is desired (szOID_ARCHIVED_KEY_ATTR)",
		"1.3.6.1.4.1.311.21.14": "szOID_CRL_SELF_CDP",
		"1.3.6.1.4.1.311.21.15": "szOID_REQUIRE_CERT_CHAIN_POLICY",
		"1.3.6.1.4.1.311.21.16": "szOID_ARCHIVED_KEY_CERT_HASH",
		"1.3.6.1.4.1.311.21.17": "szOID_ISSUED_CERT_HASH",
		"1.3.6.1.4.1.311.21.19": "szOID_DS_EMAIL_REPLICATION",
		"1.3.6.1.4.1.311.21.20": "szOID_REQUEST_CLIENT_INFO",
		"1.3.6.1.4.1.311.21.21": "szOID_ENCRYPTED_KEY_HASH",
		"1.3.6.1.4.1.311.21.22": "szOID_CERTSRV_CROSSCA_VERSION",

		"1.3.6.1.5.5.7.1.1": "Certificate Authority Information Access",

		"2.5.29.1":  "old Authority Key Identifier",
		"2.5.29.2":  "old Primary Key Attributes",
		"2.5.29.3":  "Certificate Policies",
		"2.5.29.4":  "Primary Key Usage Restriction",
		"2.5.29.9":  "Subject Directory Attributes",
		"2.5.29.14": "Subject Key Identifier",
		"2.5.29.15": "Key Usage",
		"2.5.29.16": "Private Key Usage Period",
		"2.5.29.17": "Subject Alternative Name",
		"2.5.29.18": "Issuer Alternative Name",
		"2.5.29.19": "Basic Constraints",
		"2.5.29.20": "CRL Number",
		"2.5.29.21": "Reason code",
		"2.5.29.23": "Hold Instruction Code",
		"2.5.29.24": "Invalidity Date",
		"2.5.29.27": "Delta CRL indicator",
		"2.5.29.28": "Issuing Distribution Point",
		"2.5.29.29": "Certificate Issuer",
		"2.5.29.30": "Name Constraints",
		"2.5.29.31": "CRL Distribution Points",
		"2.5.29.32": "Certificate Policies",
		"2.5.29.33": "Policy Mappings",
		"2.5.29.35": "Authority Key Identifier",
		"2.5.29.36": "Policy Constraints",
		"2.5.29.37": "Extended key usage",
		"2.5.29.46": "FreshestCRL",
		"2.5.29.54": "X.509 version 3 certificate extension Inhibit Any-policy",
	}

	mapped := mapping[oid]

	if mapped == "" {
		mapped = oid + " - UNKNOWN"
	}
	return mapped
}

//
//

func FormatCertData(certData *x509.Certificate) (formattedCertData FormattedServerCertInfo) {
	var (
		certTemplate   string
		msCertTemplate string
	)

	for _, name := range certData.Subject.Names {
		if name.Type.Equal(asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}) {
			if email, ok := name.Value.(string); ok {
				formattedCertData.SubjectEmail = email
			}
		}
	}

	formattedCertData.SubjectRDN = certData.Subject.ToRDNSequence().String()
	formattedCertData.SubjectName = certData.Subject.CommonName
	formattedCertData.RawData = certData.Raw
	formattedCertData.DNSNames = certData.DNSNames
	formattedCertData.IPAddresses = certData.IPAddresses
	formattedCertData.EmailAddresses = certData.EmailAddresses
	formattedCertData.BasicConstraintsValid = certData.BasicConstraintsValid
	formattedCertData.IssuerRDN = certData.Issuer.ToRDNSequence().String()
	formattedCertData.IssuerName = certData.Issuer.CommonName
	formattedCertData.NotBefore = certData.NotBefore.Local()
	formattedCertData.NotAfter = certData.NotAfter.Local()
	formattedCertData.CRLDistributionPoints = certData.CRLDistributionPoints
	formattedCertData.IsCA = certData.IsCA
	formattedCertData.PublicKeyAlgorithm = certData.PublicKeyAlgorithm.String()
	formattedCertData.SignatureAlgorithm = certData.SignatureAlgorithm.String()
	formattedCertData.Serial = certData.SerialNumber.String()
	thumbprint := sha1.Sum(certData.Raw)
	formattedCertData.Thumbprint = fmt.Sprintf("%x", thumbprint)
	formattedCertData.SubjectKeyId = hex.EncodeToString(certData.SubjectKeyId)
	formattedCertData.AuthorityKeyId = hex.EncodeToString(certData.AuthorityKeyId)
	formattedCertData.KeyUsage = FormatKeyUsage(certData.KeyUsage)

	for _, ext := range certData.Extensions {
		oidString := GetCertExtensionFromOID(ext.Id.String())
		if oidString == "Key Usage" ||
			oidString == "Extended key usage" ||
			oidString == "Subject Alternative Name" ||
			oidString == "Subject Key Identifier" ||
			oidString == "Authority Key Identifier" ||
			oidString == "szOID_APPLICATION_CERT_POLICIES" ||
			oidString == "CRL Distribution Points" {
			continue
		}

		if oidString == "" {
		} else if oidString == "szOID_CERTIFICATE_TEMPLATE" {
			var decoded CertTemplateExtension
			_, err := asn1.Unmarshal(ext.Value, &decoded)
			CheckFatalError(err)
			certTpl := ""
			// certTpl := GetCertTemplateNameFromOid(decoded.TemplateID.String())
			certTemplate = certTpl + " v" + strconv.Itoa(decoded.MajorVersion) + "." + strconv.Itoa(decoded.MinorVersion)
		} else if oidString == "szOID_ENROLL_CERTTYPE" {
			asn1.Unmarshal(ext.Value, &msCertTemplate)
		} else {
			formattedCertData.OtherExtensions = append(formattedCertData.OtherExtensions, oidString)
		}
	}
	if certTemplate != "" {
		formattedCertData.CertificateTemplate = certTemplate
	} else {
		formattedCertData.CertificateTemplate = msCertTemplate
	}

	if len(certData.ExtKeyUsage) > 0 {
		for _, eku := range certData.ExtKeyUsage {
			formattedCertData.ExtendedKeyUsage = append(formattedCertData.ExtendedKeyUsage, eku.String())
		}
	}

	return
}

//
//

func FormatKeyUsage(ku x509.KeyUsage) []string {
	var usages []string
	if ku&x509.KeyUsageDigitalSignature != 0 {
		usages = append(usages, "Digital Signature")
	}
	if ku&x509.KeyUsageContentCommitment != 0 {
		usages = append(usages, "Content Commitment")
	}
	if ku&x509.KeyUsageKeyEncipherment != 0 {
		usages = append(usages, "Key Encipherment")
	}
	if ku&x509.KeyUsageDataEncipherment != 0 {
		usages = append(usages, "Data Encipherment")
	}
	if ku&x509.KeyUsageKeyAgreement != 0 {
		usages = append(usages, "Key Agreement")
	}
	if ku&x509.KeyUsageCertSign != 0 {
		usages = append(usages, "Certificate Sign")
	}
	if ku&x509.KeyUsageCRLSign != 0 {
		usages = append(usages, "CRL Sign")
	}
	if ku&x509.KeyUsageEncipherOnly != 0 {
		usages = append(usages, "Encipher Only")
	}
	if ku&x509.KeyUsageDecipherOnly != 0 {
		usages = append(usages, "Decipher Only")
	}
	return usages
}

//
//

func GetCertTemplateNameFromOid(oidMappings map[string]string, oid string) string {

	mappings := make(map[string]string)

	mappedName := mappings[oid]

	if mappedName == "" {
		return oid + " - UNKNOWN"
	} else {
		return mappedName
	}
}

func GetCertTemplateOidMappings() map[string]string {
	config := GetCldConfig(nil)
	return config.CertificateManagement.CustomCertTemplateOidMappings
}
