package lib

import (
	"encoding/asn1"
	"encoding/json/v2"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ServerCertFile struct {
	TimeGenerated time.Time        `json:"TimeGenerated,omitempty,omitzero" bson:"TimeGenerated,omitempty,omitzero"`
	Certificates  []ServerCertInfo `json:"Certificates,omitempty,omitzero" bson:"Certificates,omitempty,omitzero"`
}

//
//

type CertAuthorityCertInfo struct {
	ID                             string                            `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	LastDBSync                     *time.Time                        `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	LastServerSync                 time.Time                         `json:"lastServerSync,omitempty,omitzero" bson:"lastServerSync,omitempty,omitzero"`
	CertificateAuthorityName       string                            `json:"certificateAuthorityName,omitempty,omitzero" bson:"certificateAuthorityName,omitempty,omitzero"`
	TenantName                     string                            `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	RelatedServersCertUsedOn       []ServerCertInfoServersPulledFrom `json:"relatedServersCertUsedOn,omitempty,omitzero" bson:"relatedServersCertUsedOn,omitempty,omitzero"`
	BinaryCertificate              string                            `json:"binaryCertificate,omitempty,omitzero" bson:"binaryCertificate,omitempty,omitzero"`
	BinaryPublicKey                string                            `json:"binaryPublicKey,omitempty,omitzero" bson:"binaryPublicKey,omitempty,omitzero"`
	BinaryRequest                  string                            `json:"binaryRequest,omitempty,omitzero" bson:"binaryRequest,omitempty,omitzero"`
	CallerName                     string                            `json:"callerName,omitempty,omitzero" bson:"callerName,omitempty,omitzero"`
	CertificateEffectiveDate       time.Time                         `json:"certificateEffectiveDate,omitempty,omitzero" bson:"certificateEffectiveDate,omitempty,omitzero"`
	CertificateExpirationDate      time.Time                         `json:"certificateExpirationDate,omitempty,omitzero" bson:"certificateExpirationDate,omitempty,omitzero"`
	CertificateHash                string                            `json:"certificateHash,omitempty,omitzero" bson:"certificateHash,omitempty,omitzero"`
	CertificateTemplate            string                            `json:"certificateTemplate,omitempty,omitzero" bson:"certificateTemplate,omitempty,omitzero"`
	CrlPartitionIndex              string                            `json:"crlPartitionIndex,omitempty,omitzero" bson:"crlPartitionIndex,omitempty,omitzero"`
	IssuedBinaryName               string                            `json:"issuedBinaryName,omitempty,omitzero" bson:"issuedBinaryName,omitempty,omitzero"`
	IssuedCity                     string                            `json:"issuedCity,omitempty,omitzero" bson:"issuedCity,omitempty,omitzero"`
	IssuedCommonName               string                            `json:"issuedCommonName,omitempty,omitzero" bson:"issuedCommonName,omitempty,omitzero"`
	IssuedCountryRegion            string                            `json:"issuedCountryRegion,omitempty,omitzero" bson:"issuedCountryRegion,omitempty,omitzero"`
	IssuedDistinguishedName        string                            `json:"issuedDistinguishedName,omitempty,omitzero" bson:"issuedDistinguishedName,omitempty,omitzero"`
	IssuedDomainComponent          string                            `json:"issuedDomainComponent,omitempty,omitzero" bson:"issuedDomainComponent,omitempty,omitzero"`
	IssuedEmailAddress             string                            `json:"issuedEmailAddress,omitempty,omitzero" bson:"issuedEmailAddress,omitempty,omitzero"`
	IssuedOrganization             string                            `json:"issuedOrganization,omitempty,omitzero" bson:"issuedOrganization,omitempty,omitzero"`
	IssuedOrganizationUnit         string                            `json:"issuedOrganizationUnit,omitempty,omitzero" bson:"issuedOrganizationUnit,omitempty,omitzero"`
	IssuedRequestID                string                            `json:"issuedRequestId,omitempty,omitzero" bson:"issuedRequestId,omitempty,omitzero"`
	IssuedState                    string                            `json:"issuedState,omitempty,omitzero" bson:"issuedState,omitempty,omitzero"`
	IssuedSubjectKeyIdentifier     string                            `json:"issuedSubjectKeyIdentifier,omitempty,omitzero" bson:"issuedSubjectKeyIdentifier,omitempty,omitzero"`
	IssuerNameID                   string                            `json:"issuerNameId,omitempty,omitzero" bson:"issuerNameId,omitempty,omitzero"`
	OldCertificate                 string                            `json:"oldCertificate,omitempty,omitzero" bson:"oldCertificate,omitempty,omitzero"`
	PublicKeyAlgorithm             string                            `json:"publicKeyAlgorithm,omitempty,omitzero" bson:"publicKeyAlgorithm,omitempty,omitzero"`
	PublicKeyAlgorithmParameters   string                            `json:"publicKeyAlgorithmParameters,omitempty,omitzero" bson:"publicKeyAlgorithmParameters,omitempty,omitzero"`
	PublicKeyLength                string                            `json:"publicKeyLength,omitempty,omitzero" bson:"publicKeyLength,omitempty,omitzero"`
	PublishExpiredCertificateInCrl bool                              `json:"publishExpiredCertificateInCrl,omitempty,omitzero" bson:"publishExpiredCertificateInCrl,omitempty,omitzero"`
	RequestAttributes              string                            `json:"requestAttributes,omitempty,omitzero" bson:"requestAttributes,omitempty,omitzero"`
	RequestBinaryName              string                            `json:"requestBinaryName,omitempty,omitzero" bson:"requestBinaryName,omitempty,omitzero"`
	RequestCity                    string                            `json:"requestCity,omitempty,omitzero" bson:"requestCity,omitempty,omitzero"`
	RequestCommonName              string                            `json:"requestCommonName,omitempty,omitzero" bson:"requestCommonName,omitempty,omitzero"`
	RequestCountryRegion           string                            `json:"requestCountryRegion,omitempty,omitzero" bson:"requestCountryRegion,omitempty,omitzero"`
	RequestDisposition             string                            `json:"requestDisposition,omitempty,omitzero" bson:"requestDisposition,omitempty,omitzero"`
	RequestDispositionMessage      string                            `json:"requestDispositionMessage,omitempty,omitzero" bson:"requestDispositionMessage,omitempty,omitzero"`
	RequestDistinguishedName       string                            `json:"requestDistinguishedName,omitempty,omitzero" bson:"requestDistinguishedName,omitempty,omitzero"`
	RequestDomainComponent         string                            `json:"requestDomainComponent,omitempty,omitzero" bson:"requestDomainComponent,omitempty,omitzero"`
	RequestFlags                   string                            `json:"requestFlags,omitempty,omitzero" bson:"requestFlags,omitempty,omitzero"`
	RequestID                      string                            `json:"requestId,omitempty,omitzero" bson:"requestId,omitempty,omitzero"`
	RequestOrganization            string                            `json:"requestOrganization,omitempty,omitzero" bson:"requestOrganization,omitempty,omitzero"`
	RequestOrganizationUnit        string                            `json:"requestOrganizationUnit,omitempty,omitzero" bson:"requestOrganizationUnit,omitempty,omitzero"`
	RequestResolutionDate          time.Time                         `json:"requestResolutionDate,omitempty,omitzero" bson:"requestResolutionDate,omitempty,omitzero"`
	RequestState                   string                            `json:"requestState,omitempty,omitzero" bson:"requestState,omitempty,omitzero"`
	RequestStatusCode              string                            `json:"requestStatusCode,omitempty,omitzero" bson:"requestStatusCode,omitempty,omitzero"`
	RequestSubmissionDate          time.Time                         `json:"requestSubmissionDate,omitempty,omitzero" bson:"requestSubmissionDate,omitempty,omitzero"`
	RequestType                    string                            `json:"requestType,omitempty,omitzero" bson:"requestType,omitempty,omitzero"`
	RequesterName                  string                            `json:"requesterName,omitempty,omitzero" bson:"requesterName,omitempty,omitzero"`
	SerialNumber                   string                            `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	SignerApplicationPolicies      string                            `json:"signerApplicationPolicies,omitempty,omitzero" bson:"signerApplicationPolicies,omitempty,omitzero"`
	SignerPolicies                 string                            `json:"signerPolicies,omitempty,omitzero" bson:"signerPolicies,omitempty,omitzero"`
	TemplateEnrollmentFlags        string                            `json:"templateEnrollmentFlags,omitempty,omitzero" bson:"templateEnrollmentFlags,omitempty,omitzero"`
	TemplateGeneralFlags           string                            `json:"templateGeneralFlags,omitempty,omitzero" bson:"templateGeneralFlags,omitempty,omitzero"`
	TemplatePrivateKeyFlags        string                            `json:"templatePrivateKeyFlags,omitempty,omitzero" bson:"templatePrivateKeyFlags,omitempty,omitzero"`
	UserPrincipalName              string                            `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}

//
//

type ServerCertInfo struct {
	ID                       string                            `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	IgnoreExpiration         bool                              `json:"ignoreExpiration,omitempty,omitzero" bson:"ignoreExpiration,omitempty,omitzero"`
	LastDBSync               *time.Time                        `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	ServerSyncTime           UnixTimeFromPowershell            `json:"serverSyncTime,omitempty,omitzero" bson:"serverSyncTime,omitempty,omitzero"`
	PulledFromServer         *string                           `json:"pulledFromServer,omitempty,omitzero" bson:"pulledFromServer,omitempty,omitzero"`
	ServersPulledFrom        []ServerCertInfoServersPulledFrom `json:"serversPulledFrom,omitempty,omitzero" bson:"serversPulledFrom,omitempty,omitzero"`
	TenantName               *string                           `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	TenantNames              []string                          `json:"tenantNames,omitempty,omitzero" bson:"tenantNames,omitempty,omitzero"`
	RelatedCertAuthData      *CertAuthorityCertInfo            `json:"relatedCertAuthData,omitempty,omitzero" bson:"relatedCertAuthData,omitempty,omitzero"`
	Archived                 bool                              `json:"archived,omitempty,omitzero" bson:"archived,omitempty,omitzero"`
	DnsNameList              *[]string                         `json:"dnsNameList,omitempty,omitzero" bson:"dnsNameList,omitempty,omitzero"`
	DNSNames                 []string                          `json:"dnsNames,omitempty,omitzero" bson:"dnsNames,omitempty,omitzero"`
	EmailAddresses           []string                          `json:"emailAddresses,omitempty,omitzero" bson:"emailAddresses,omitempty,omitzero"`
	IPAddresses              []net.IP                          `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	URIs                     []*url.URL                        `json:"uris,omitempty,omitzero" bson:"uris,omitempty,omitzero"`
	EnhancedKeyUsageList     *[]string                         `json:"enhancedKeyUsageList,omitempty,omitzero" bson:"enhancedKeyUsageList,omitempty,omitzero"`
	EnrollmentPolicyEndPoint *struct {
		AuthenticationType float64 `json:"authenticationType,omitempty,omitzero" bson:"authenticationType,omitempty,omitzero"`
		URL                *string `json:"url,omitempty,omitzero" bson:"url,omitempty,omitzero"`
	} `json:"EnrollmentPolicyEndPoint,omitempty,omitzero" bson:"EnrollmentPolicyEndPoint,omitempty,omitzero"`
	EnrollmentServerEndPoint *struct {
		AuthenticationType float64 `json:"authenticationType,omitempty,omitzero" bson:"authenticationType,omitempty,omitzero"`
		URL                *string `json:"url,omitempty,omitzero" bson:"url,omitempty,omitzero"`
	} `json:"enrollmentServerEndPoint,omitempty,omitzero" bson:"enrollmentServerEndPoint,omitempty,omitzero"`
	Extensions            *[]string `json:"extensions,omitempty,omitzero" bson:"extensions,omitempty,omitzero"`
	FriendlyName          string    `json:"friendlyName,omitempty,omitzero" bson:"friendlyName,omitempty,omitzero"`
	Handle                float64   `json:"handle,omitempty,omitzero" bson:"handle,omitempty,omitzero"`
	HasPrivateKey         bool      `json:"hasPrivateKey,omitempty,omitzero" bson:"hasPrivateKey,omitempty,omitzero"`
	Issuer                string    `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
	IssuingCertificateURL []string  `json:"issuingCertificateURL,omitempty,omitzero" bson:"issuingCertificateURL,omitempty,omitzero"`
	IssuerName            *struct {
		Name    string  `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Oid     string  `json:"oid,omitempty,omitzero" bson:"oid,omitempty,omitzero"`
		RawData *string `json:"rawData,omitempty,omitzero" bson:"rawData,omitempty,omitzero"`
	} `json:"issuerName,omitempty,omitzero" bson:"issuerName,omitempty,omitzero"`
	NotAfter  time.Time `json:"notAfter,omitempty,omitzero" bson:"notAfter,omitempty,omitzero"`
	NotBefore time.Time `json:"notBefore,omitempty,omitzero" bson:"notBefore,omitempty,omitzero"`
	PolicyID  *string   `json:"policyId,omitempty,omitzero" bson:"policyId,omitempty,omitzero"`
	// PrivateKey *struct {
	// 	CspKeyContainerInfo  string  `json:"cspKeyContainerInfo,omitempty,omitzero" bson:"cspKeyContainerInfo,omitempty,omitzero"`
	// 	KeyExchangeAlgorithm string  `json:"keyExchangeAlgorithm,omitempty,omitzero" bson:"keyExchangeAlgorithm,omitempty,omitzero"`
	// 	KeySize              float64 `json:"keySize,omitempty,omitzero" bson:"keySize,omitempty,omitzero"`
	// 	LegalKeySizes        string  `json:"legalKeySizes,omitempty,omitzero" bson:"legalKeySizes,omitempty,omitzero"`
	// 	PersistKeyInCsp      bool    `json:"persistKeyInCsp,omitempty,omitzero" bson:"persistKeyInCsp,omitempty,omitzero"`
	// 	PublicOnly           bool    `json:"publicOnly,omitempty,omitzero" bson:"publicOnly,omitempty,omitzero"`
	// 	SignatureAlgorithm   string  `json:"signatureAlgorithm,omitempty,omitzero" bson:"signatureAlgorithm,omitempty,omitzero"`
	// } `json:"privateKey,omitempty,omitzero" bson:"privateKey,omitempty,omitzero"`
	PublicKey struct {
		EncodedKeyValue   string  `json:"encodedKeyValue,omitempty,omitzero" bson:"encodedKeyValue,omitempty,omitzero"`
		EncodedParameters string  `json:"encodedParameters,omitempty,omitzero" bson:"encodedParameters,omitempty,omitzero"`
		Key               *string `json:"key,omitempty,omitzero" bson:"key,omitempty,omitzero"`
		Oid               string  `json:"oid,omitempty,omitzero" bson:"oid,omitempty,omitzero"`
	} `json:"publicKey,omitempty,omitzero" bson:"publicKey,omitempty,omitzero"`
	RawData             []byte `json:"rawData,omitempty,omitzero" bson:"rawData,omitempty,omitzero"`
	SendAsTrustedIssuer bool   `json:"sendAsTrustedIssuer,omitempty,omitzero" bson:"sendAsTrustedIssuer,omitempty,omitzero"`
	SerialNumber        string `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	SignatureAlgorithm  struct {
		FriendlyName string `json:"friendlyName,omitempty,omitzero" bson:"friendlyName,omitempty,omitzero"`
		Value        string `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
	} `json:"signatureAlgorithm,omitempty,omitzero" bson:"signatureAlgorithm,omitempty,omitzero"`
	Subject     string `json:"subject,omitempty,omitzero" bson:"subject,omitempty,omitzero"`
	SubjectName struct {
		Name    string  `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Oid     string  `json:"oid,omitempty,omitzero" bson:"oid,omitempty,omitzero"`
		RawData *string `json:"rawData,omitempty,omitzero" bson:"rawData,omitempty,omitzero"`
	} `json:"subjectName,omitempty,omitzero" bson:"subjectName,omitempty,omitzero"`
	Thumbprint string  `json:"thumbprint,omitempty,omitzero" bson:"thumbprint,omitempty,omitzero"`
	Version    float64 `json:"version,omitempty,omitzero" bson:"version,omitempty,omitzero"`
	ParentPath *string `json:"parentPath,omitempty,omitzero" bson:"parentPath,omitempty,omitzero"`
}

type UnixTimeFromPowershell time.Time

func (ut *UnixTimeFromPowershell) UnmarshalJSON(b []byte) error {
	// Parse the JSON number as an int64
	// var timestamp int64
	// if err := json.Unmarshal(strconv.Atoi(string(b)), &timestamp); err != nil {
	// 	return err
	// }
	var tsStr string
	err := json.Unmarshal(b, &tsStr)
	if err != nil {
		fmt.Println("err := json.Unmarshal(b, &tsStr)")
		return err
	}
	tsStrSpl := strings.Split(tsStr, ".")[0]
	timestamp, err := strconv.Atoi(tsStrSpl)
	if err != nil {
		fmt.Println("timestamp, err := strconv.Atoi(tsStrSpl)")
		return err
	}
	// Update the UnixTime value
	*ut = UnixTimeFromPowershell(time.Unix(int64(timestamp), 0))
	return nil
}

//
//

type PowershellDateObject time.Time

func (ut *PowershellDateObject) UnmarshalJSON(b []byte) error {
	// Parse the JSON number as an int64
	// var timestamp int64
	// if err := json.Unmarshal(strconv.Atoi(string(b)), &timestamp); err != nil {
	// 	return err
	// }
	// var tsStr string
	// err := json.Unmarshal(b, &tsStr)
	// if err != nil {
	// 	return err
	// }
	// tsStrSpl := strings.Split(tsStr, ".")[0]
	// timestamp, err := strconv.Atoi(tsStrSpl)
	// if err != nil {
	// 	return err
	// }

	var pdoStr string
	err := json.Unmarshal(b, &pdoStr)

	valStr := strings.TrimPrefix(pdoStr, "/Date(")
	valStr = strings.TrimSuffix(valStr, ")/")
	valInt, err := strconv.ParseInt(valStr, 10, 64)
	CheckFatalError(err)
	valDate := time.UnixMilli(valInt)

	// Update the UnixTime value
	*ut = PowershellDateObject(valDate)
	return nil
}

//
//

type ServerCertInfoRaw struct {
	// Archived                 bool     `json:"Archived"`
	DnsNameList          []string `json:"DnsNameList"`
	EnhancedKeyUsageList []string `json:"EnhancedKeyUsageList"`
	// EnrollmentPolicyEndPoint struct {
	// 	AuthenticationType float64 `json:"AuthenticationType"`
	// 	URL                any     `json:"Url"`
	// } `json:"EnrollmentPolicyEndPoint"`
	// EnrollmentServerEndPoint struct {
	// 	AuthenticationType float64 `json:"AuthenticationType"`
	// 	URL                any     `json:"Url"`
	// } `json:"EnrollmentServerEndPoint"`
	// Extensions    []string `json:"Extensions"`
	FriendlyName string `json:"FriendlyName"`
	// Handle        float64  `json:"Handle"`
	HasPrivateKey bool   `json:"HasPrivateKey"`
	Issuer        string `json:"Issuer"`
	// IssuerName    struct {
	// 	Name    string `json:"Name"`
	// 	Oid     string `json:"Oid"`
	// 	RawData string `json:"RawData"`
	// } `json:"IssuerName"`
	NotAfter  PowershellDateObject `json:"NotAfter"`
	NotBefore PowershellDateObject `json:"NotBefore"`
	// PsChildName string `json:"PSChildName"`
	// PsDrive     struct {
	// 	Credential      string `json:"Credential"`
	// 	CurrentLocation string `json:"CurrentLocation"`
	// 	Description     string `json:"Description"`
	// 	DisplayRoot     any    `json:"DisplayRoot"`
	// 	MaximumSize     any    `json:"MaximumSize"`
	// 	Name            string `json:"Name"`
	// 	Provider        string `json:"Provider"`
	// 	Root            string `json:"Root"`
	// } `json:"PSDrive"`
	// PsIsContainer bool   `json:"PSIsContainer"`
	PsParentPath string `json:"PSParentPath"`
	// PsPath        string `json:"PSPath"`
	// PsProvider    struct {
	// 	Capabilities     float64 `json:"Capabilities"`
	// 	Description      string  `json:"Description"`
	// 	Drives           string  `json:"Drives"`
	// 	HelpFile         string  `json:"HelpFile"`
	// 	Home             string  `json:"Home"`
	// 	ImplementingType string  `json:"ImplementingType"`
	// 	Module           string  `json:"Module"`
	// 	ModuleName       string  `json:"ModuleName"`
	// 	Name             string  `json:"Name"`
	// 	PsSnapIn         any     `json:"PSSnapIn"`
	// } `json:"PSProvider"`
	// PolicyID   any `json:"PolicyId"`
	// PrivateKey *struct {
	// 	CspKeyContainerInfo  string  `json:"CspKeyContainerInfo"`
	// 	KeyExchangeAlgorithm string  `json:"KeyExchangeAlgorithm"`
	// 	KeySize              float64 `json:"KeySize"`
	// 	LegalKeySizes        string  `json:"LegalKeySizes"`
	// 	PersistKeyInCsp      bool    `json:"PersistKeyInCsp"`
	// 	PublicOnly           bool    `json:"PublicOnly"`
	// 	SignatureAlgorithm   string  `json:"SignatureAlgorithm"`
	// } `json:"PrivateKey"`
	// PublicKey struct {
	// 	EncodedKeyValue   string  `json:"EncodedKeyValue"`
	// 	EncodedParameters string  `json:"EncodedParameters"`
	// 	Key               *string `json:"Key"`
	// 	Oid               string  `json:"Oid"`
	// } `json:"PublicKey"`
	RawData []byte `json:"RawData"`
	// SendAsTrustedIssuer bool   `json:"SendAsTrustedIssuer"`
	SerialNumber string `json:"SerialNumber"`
	// ServerSyncTime string `json:"ServerSyncTime"`
	ServerSyncTime     UnixTimeFromPowershell `json:"ServerSyncTime"`
	SignatureAlgorithm struct {
		FriendlyName string `json:"FriendlyName"`
		Value        string `json:"Value"`
	} `json:"SignatureAlgorithm"`
	Subject string `json:"Subject"`
	// SubjectName struct {
	// 	Name    string `json:"Name"`
	// 	Oid     string `json:"Oid"`
	// 	RawData string `json:"RawData"`
	// } `json:"SubjectName"`
	Thumbprint string `json:"Thumbprint"`
	// Version    float64 `json:"Version"`
}

// type ServerCertInfoFile struct {
// 	Certificates
// }

//
//

type FormattedServerCertInfo struct {
	IssuerName            string                            `json:"issuerName,omitempty,omitzero" bson:"issuerName,omitempty,omitzero"`
	IssuerRDN             string                            `json:"issuerRDN,omitempty,omitzero" bson:"issuerRDN,omitempty,omitzero"`
	DNSNames              []string                          `json:"dnsNames,omitempty,omitzero" bson:"dnsNames,omitempty,omitzero"`
	EmailAddresses        []string                          `json:"emailAddresses,omitempty,omitzero" bson:"emailAddresses,omitempty,omitzero"`
	IPAddresses           []net.IP                          `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	SubjectName           string                            `json:"subjectName,omitempty,omitzero" bson:"subjectName,omitempty,omitzero"`
	SubjectRDN            string                            `json:"subjectRDN,omitempty,omitzero" bson:"subjectRDN,omitempty,omitzero"`
	SubjectEmail          string                            `json:"subjectEmail,omitempty,omitzero" bson:"subjectEmail,omitempty,omitzero"`
	CRLDistributionPoints []string                          `json:"crlDistributionPoints,omitempty,omitzero" bson:"crlDistributionPoints,omitempty,omitzero"`
	CertificateTemplate   string                            `json:"certificateTemplate,omitempty,omitzero" bson:"certificateTemplate,omitempty,omitzero"`
	RawData               []byte                            `json:"rawData,omitempty,omitzero" bson:"rawData,omitempty,omitzero"`
	PublicKeyAlgorithm    string                            `json:"publicKeyAlgorithm,omitempty,omitzero" bson:"publicKeyAlgorithm,omitempty,omitzero"`
	SignatureAlgorithm    string                            `json:"signatureAlgorithm,omitempty,omitzero" bson:"signatureAlgorithm,omitempty,omitzero"`
	IsCA                  bool                              `json:"isCA,omitempty,omitzero" bson:"isCA,omitempty,omitzero"`
	Serial                string                            `json:"serial,omitempty,omitzero" bson:"serial,omitempty,omitzero"`
	SerialWindows         string                            `json:"serialWindows,omitempty,omitzero" bson:"serialWindows,omitempty,omitzero"`
	SubjectKeyId          string                            `json:"subjectKeyId,omitempty,omitzero" bson:"subjectKeyId,omitempty,omitzero"`
	AuthorityKeyId        string                            `json:"authorityKeyId,omitempty,omitzero" bson:"authorityKeyId,omitempty,omitzero"`
	ExtendedKeyUsage      []string                          `json:"extendedKeyUsage,omitempty,omitzero" bson:"extendedKeyUsage,omitempty,omitzero"`
	KeyUsage              []string                          `json:"keyUsage,omitempty,omitzero" bson:"keyUsage,omitempty,omitzero"`
	OtherExtensions       []string                          `json:"otherExtensions,omitempty,omitzero" bson:"otherExtensions,omitempty,omitzero"`
	NotBefore             time.Time                         `json:"notBefore,omitempty,omitzero" bson:"notBefore,omitempty,omitzero"`
	NotAfter              time.Time                         `json:"notAfter,omitempty,omitzero" bson:"notAfter,omitempty,omitzero"`
	BasicConstraintsValid bool                              `json:"basicConstraintsValid,omitempty,omitzero" bson:"basicConstraintsValid,omitempty,omitzero"`
	Thumbprint            string                            `json:"thumbprint,omitempty,omitzero" bson:"thumbprint,omitempty,omitzero"`
	ParentPath            *string                           `json:"parentPath,omitempty,omitzero" bson:"parentPath,omitempty,omitzero"`
	RelatedCertAuthData   *CertAuthorityCertInfo            `json:"relatedCertAuthData,omitempty,omitzero" bson:"relatedCertAuthData,omitempty,omitzero"`
	Id                    string                            `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	LastDBSync            *time.Time                        `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	ServerSyncTime        time.Time                         `json:"serverSyncTime,omitempty,omitzero" bson:"serverSyncTime,omitempty,omitzero"`
	PulledFromServer      *string                           `json:"pulledFromServer,omitempty,omitzero" bson:"pulledFromServer,omitempty,omitzero"`
	ServersPulledFrom     []ServerCertInfoServersPulledFrom `json:"serversPulledFrom,omitempty,omitzero" bson:"serversPulledFrom,omitempty,omitzero"`
	TenantName            *string                           `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	TenantNames           []string                          `json:"tenantNames,omitempty,omitzero" bson:"tenantNames,omitempty,omitzero"`
	FriendlyNames         map[string][]string               `json:"friendlyNames,omitempty,omitzero" bson:"friendlyNames,omitempty,omitzero"`
	FriendlyName          *string                           `json:"friendlyName,omitempty,omitzero" bson:"friendlyName,omitempty,omitzero"`
	HasPrivateKey         bool                              `json:"hasPrivateKey,omitempty,omitzero" bson:"hasPrivateKey,omitempty,omitzero"`
	ServersWithPrivateKey *[]string                         `json:"serversWithPrivateKey,omitempty,omitzero" bson:"serversWithPrivateKey,omitempty,omitzero"`
}

type FriendlyName struct {
	Name    string   `json:"name"`
	Servers []string `json:"servers"`
}

type HasPrivateKey struct {
}

//
//

type CertTemplateExtension struct {
	TemplateID   asn1.ObjectIdentifier
	MajorVersion int `asn1:"optional"`
	MinorVersion int `asn1:"optional"`
}

//
//

// type ServerCertInfoServersPulledFrom map[string][]string

type ServerCertInfoServersPulledFrom struct {
	ServerName       string   `json:"serverName,omitempty,omitzero" bson:"serverName,omitempty,omitzero"`
	CertificatePaths []string `json:"certificatePaths,omitempty,omitzero" bson:"certificatePaths,omitempty,omitzero"`
}

//
//
