package lib

import (
	"net"
	"net/url"
	"time"
)

type PackerLogBuildData struct {
	AzDoBuildName          string                                     `json:"azDoBuildName,omitempty" bson:"azDoBuildName,omitempty"`
	AzDoCompleteTime       time.Time                                  `json:"azDoCompleteTime,omitempty" bson:"azDoCompleteTime,omitempty"`
	AzDoDuration           time.Duration                              `json:"azDoDuration,omitempty" bson:"azDoDuration,omitempty"`
	AzDoHost               string                                     `json:"azDoHost,omitempty" bson:"azDoHost,omitempty"`
	AzDoLogFile            string                                     `json:"azDoLogFile,omitempty" bson:"azDoLogFile,omitempty"`
	AzDoStartTime          time.Time                                  `json:"azDoStartTime,omitempty" bson:"azDoStartTime,omitempty"`
	BuildBaseImageVersion  *AzureResourceStorageProfileImageReference `json:"buildBaseImageVersion,omitempty" bson:"buildBaseImageVersion,omitempty"`
	BuildImageCompleteTime time.Time                                  `json:"buildImageCompleteTime,omitempty" bson:"buildImageCompleteTime,omitempty"`
	BuildImageDuration     time.Duration                              `json:"buildImageDuration,omitempty" bson:"buildImageDuration,omitempty"`
	BuildImageEnv          string                                     `json:"buildImageEnv,omitempty" bson:"buildImageEnv,omitempty"`
	BuildImageStartTime    time.Time                                  `json:"buildImageStartTime,omitempty" bson:"buildImageStartTime,omitempty"`
	OutputImgDef           string                                     `json:"outputImgDef,omitempty" bson:"outputImgDef,omitempty"`
	OutputImgGalleryName   string                                     `json:"outputImgGalleryName,omitempty" bson:"outputImgGalleryName,omitempty"`
	OutputImgId            string                                     `json:"outputImgId,omitempty" bson:"outputImgId,omitempty"`
	OutputImgResGrp        string                                     `json:"outputImgResGrp,omitempty" bson:"outputImgResGrp,omitempty"`
	OutputImgVersion       string                                     `json:"outputImgVersion,omitempty" bson:"outputImgVersion,omitempty"`
}

//
//

type PackerPublishImageResponse struct {
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Location   string `json:"location,omitempty" bson:"location,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublishingProfile struct {
			ExcludeFromLatest  bool      `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
			PublishedDate      time.Time `json:"publishedDate,omitempty" bson:"publishedDate,omitempty"`
			ReplicaCount       float64   `json:"replicaCount,omitempty" bson:"replicaCount,omitempty"`
			ReplicationMode    string    `json:"replicationMode,omitempty" bson:"replicationMode,omitempty"`
			StorageAccountType string    `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			TargetRegions      []struct {
				Name                 string  `json:"name,omitempty" bson:"name,omitempty"`
				RegionalReplicaCount float64 `json:"regionalReplicaCount,omitempty" bson:"regionalReplicaCount,omitempty"`
				StorageAccountType   string  `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			} `json:"targetRegions,omitempty" bson:"targetRegions,omitempty"`
		} `json:"publishingProfile,omitempty" bson:"publishingProfile,omitempty"`
		SafetyProfile struct {
			AllowDeletionOfReplicatedLocations bool `json:"allowDeletionOfReplicatedLocations,omitempty" bson:"allowDeletionOfReplicatedLocations,omitempty"`
			ReportedForPolicyViolation         bool `json:"reportedForPolicyViolation,omitempty" bson:"reportedForPolicyViolation,omitempty"`
		} `json:"safetyProfile,omitempty" bson:"safetyProfile,omitempty"`
		StorageProfile struct {
			Source struct {
				VirtualMachineID string `json:"virtualMachineId,omitempty" bson:"virtualMachineId,omitempty"`
			} `json:"source,omitempty" bson:"source,omitempty"`
		} `json:"storageProfile,omitempty" bson:"storageProfile,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Tags map[string]string `json:"tags" bson:"tags"`
	Type string            `json:"type,omitempty" bson:"type,omitempty"`
}

//
//

type CertAuthorityCertInfo struct {
	ID                             string                            `json:"id,omitempty" bson:"_id,omitempty"`
	LastDBSync                     *time.Time                        `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
	LastServerSync                 time.Time                         `json:"lastServerSync,omitempty"  bson:"lastServerSync,omitempty"`
	CertificateAuthorityName       string                            `json:"certificateAuthorityName,omitempty" bson:"certificateAuthorityName,omitempty"`
	TenantName                     string                            `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	RelatedServersCertUsedOn       []ServerCertInfoServersPulledFrom `json:"relatedServersCertUsedOn,omitempty" bson:"relatedServersCertUsedOn,omitempty"`
	BinaryCertificate              string                            `json:"binaryCertificate,omitempty" bson:"binaryCertificate,omitempty"`
	BinaryPublicKey                string                            `json:"binaryPublicKey,omitempty" bson:"binaryPublicKey,omitempty"`
	BinaryRequest                  string                            `json:"binaryRequest,omitempty" bson:"binaryRequest,omitempty"`
	CallerName                     string                            `json:"callerName,omitempty" bson:"callerName,omitempty"`
	CertificateEffectiveDate       time.Time                         `json:"certificateEffectiveDate,omitempty" bson:"certificateEffectiveDate,omitempty"`
	CertificateExpirationDate      time.Time                         `json:"certificateExpirationDate,omitempty" bson:"certificateExpirationDate,omitempty"`
	CertificateHash                string                            `json:"certificateHash,omitempty" bson:"certificateHash,omitempty"`
	CertificateTemplate            string                            `json:"certificateTemplate,omitempty" bson:"certificateTemplate,omitempty"`
	CrlPartitionIndex              string                            `json:"crlPartitionIndex,omitempty" bson:"crlPartitionIndex,omitempty"`
	IssuedBinaryName               string                            `json:"issuedBinaryName,omitempty" bson:"issuedBinaryName,omitempty"`
	IssuedCity                     string                            `json:"issuedCity,omitempty" bson:"issuedCity,omitempty"`
	IssuedCommonName               string                            `json:"issuedCommonName,omitempty" bson:"issuedCommonName,omitempty"`
	IssuedCountryRegion            string                            `json:"issuedCountryRegion,omitempty" bson:"issuedCountryRegion,omitempty"`
	IssuedDistinguishedName        string                            `json:"issuedDistinguishedName,omitempty" bson:"issuedDistinguishedName,omitempty"`
	IssuedDomainComponent          string                            `json:"issuedDomainComponent,omitempty" bson:"issuedDomainComponent,omitempty"`
	IssuedEmailAddress             string                            `json:"issuedEmailAddress,omitempty" bson:"issuedEmailAddress,omitempty"`
	IssuedOrganization             string                            `json:"issuedOrganization,omitempty" bson:"issuedOrganization,omitempty"`
	IssuedOrganizationUnit         string                            `json:"issuedOrganizationUnit,omitempty" bson:"issuedOrganizationUnit,omitempty"`
	IssuedRequestID                string                            `json:"issuedRequestId,omitempty" bson:"issuedRequestId,omitempty"`
	IssuedState                    string                            `json:"issuedState,omitempty" bson:"issuedState,omitempty"`
	IssuedSubjectKeyIdentifier     string                            `json:"issuedSubjectKeyIdentifier,omitempty" bson:"issuedSubjectKeyIdentifier,omitempty"`
	IssuerNameID                   string                            `json:"issuerNameId,omitempty" bson:"issuerNameId,omitempty"`
	OldCertificate                 string                            `json:"oldCertificate,omitempty" bson:"oldCertificate,omitempty"`
	PublicKeyAlgorithm             string                            `json:"publicKeyAlgorithm,omitempty" bson:"publicKeyAlgorithm,omitempty"`
	PublicKeyAlgorithmParameters   string                            `json:"publicKeyAlgorithmParameters,omitempty" bson:"publicKeyAlgorithmParameters,omitempty"`
	PublicKeyLength                string                            `json:"publicKeyLength,omitempty" bson:"publicKeyLength,omitempty"`
	PublishExpiredCertificateInCrl bool                              `json:"publishExpiredCertificateInCrl,omitempty" bson:"publishExpiredCertificateInCrl,omitempty"`
	RequestAttributes              string                            `json:"requestAttributes,omitempty" bson:"requestAttributes,omitempty"`
	RequestBinaryName              string                            `json:"requestBinaryName,omitempty" bson:"requestBinaryName,omitempty"`
	RequestCity                    string                            `json:"requestCity,omitempty" bson:"requestCity,omitempty"`
	RequestCommonName              string                            `json:"requestCommonName,omitempty" bson:"requestCommonName,omitempty"`
	RequestCountryRegion           string                            `json:"requestCountryRegion,omitempty" bson:"requestCountryRegion,omitempty"`
	RequestDisposition             string                            `json:"requestDisposition,omitempty" bson:"requestDisposition,omitempty"`
	RequestDispositionMessage      string                            `json:"requestDispositionMessage,omitempty" bson:"requestDispositionMessage,omitempty"`
	RequestDistinguishedName       string                            `json:"requestDistinguishedName,omitempty" bson:"requestDistinguishedName,omitempty"`
	RequestDomainComponent         string                            `json:"requestDomainComponent,omitempty" bson:"requestDomainComponent,omitempty"`
	RequestFlags                   string                            `json:"requestFlags,omitempty" bson:"requestFlags,omitempty"`
	RequestID                      string                            `json:"requestId,omitempty" bson:"requestId,omitempty"`
	RequestOrganization            string                            `json:"requestOrganization,omitempty" bson:"requestOrganization,omitempty"`
	RequestOrganizationUnit        string                            `json:"requestOrganizationUnit,omitempty" bson:"requestOrganizationUnit,omitempty"`
	RequestResolutionDate          time.Time                         `json:"requestResolutionDate,omitempty" bson:"requestResolutionDate,omitempty"`
	RequestState                   string                            `json:"requestState,omitempty" bson:"requestState,omitempty"`
	RequestStatusCode              string                            `json:"requestStatusCode,omitempty" bson:"requestStatusCode,omitempty"`
	RequestSubmissionDate          time.Time                         `json:"requestSubmissionDate,omitempty" bson:"requestSubmissionDate,omitempty"`
	RequestType                    string                            `json:"requestType,omitempty" bson:"requestType,omitempty"`
	RequesterName                  string                            `json:"requesterName,omitempty" bson:"requesterName,omitempty"`
	SerialNumber                   string                            `json:"serialNumber,omitempty" bson:"serialNumber,omitempty"`
	SignerApplicationPolicies      string                            `json:"signerApplicationPolicies,omitempty" bson:"signerApplicationPolicies,omitempty"`
	SignerPolicies                 string                            `json:"signerPolicies,omitempty" bson:"signerPolicies,omitempty"`
	TemplateEnrollmentFlags        string                            `json:"templateEnrollmentFlags,omitempty" bson:"templateEnrollmentFlags,omitempty"`
	TemplateGeneralFlags           string                            `json:"templateGeneralFlags,omitempty" bson:"templateGeneralFlags,omitempty"`
	TemplatePrivateKeyFlags        string                            `json:"templatePrivateKeyFlags,omitempty" bson:"templatePrivateKeyFlags,omitempty"`
	UserPrincipalName              string                            `json:"userPrincipalName,omitempty" bson:"userPrincipalName,omitempty"`
}

//
//

type ServerCertInfo struct {
	ID                       string                            `json:"id,omitempty" bson:"_id,omitempty"`
	IgnoreExpiration         bool                              `json:"ignoreExpiration,omitempty"  bson:"ignoreExpiration"`
	LastDBSync               *time.Time                        `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
	LastServerSync           time.Time                         `json:"lastServerSync,omitempty"  bson:"lastServerSync,omitempty"`
	PulledFromServer         *string                           `json:"pulledFromServer,omitempty" bson:"pulledFromServer,omitempty"`
	ServersPulledFrom        []ServerCertInfoServersPulledFrom `json:"serversPulledFrom,omitempty" bson:"serversPulledFrom,omitempty"`
	TenantName               *string                           `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	TenantNames              []string                          `json:"tenantNames,omitempty" bson:"tenantNames,omitempty"`
	RelatedCertAuthData      *CertAuthorityCertInfo            `json:"relatedCertAuthData,omitempty" bson:"relatedCertAuthData,omitempty"`
	Archived                 bool                              `json:"archived,omitempty" bson:"archived,omitempty"`
	DnsNameList              *[]string                         `json:"dnsNameList,omitempty" bson:"dnsNameList,omitempty"`
	DNSNames                 []string                          `json:"dnsNames,omitempty,omitzero" bson:"dnsNames,omitempty,omitzero"`
	EmailAddresses           []string                          `json:"emailAddresses,omitempty,omitzero" bson:"emailAddresses,omitempty,omitzero"`
	IPAddresses              []net.IP                          `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	URIs                     []*url.URL                        `json:"uris,omitempty,omitzero" bson:"uris,omitempty,omitzero"`
	EnhancedKeyUsageList     *[]string                         `json:"enhancedKeyUsageList,omitempty" bson:"enhancedKeyUsageList,omitempty"`
	EnrollmentPolicyEndPoint *struct {
		AuthenticationType float64 `json:"authenticationType,omitempty" bson:"authenticationType,omitempty"`
		URL                *string `json:"url,omitempty" bson:"url,omitempty"`
	} `json:"EnrollmentPolicyEndPoint,omitempty" bson:"EnrollmentPolicyEndPoint,omitempty"`
	EnrollmentServerEndPoint *struct {
		AuthenticationType float64 `json:"authenticationType,omitempty" bson:"authenticationType,omitempty"`
		URL                *string `json:"url,omitempty" bson:"url,omitempty"`
	} `json:"enrollmentServerEndPoint,omitempty" bson:"enrollmentServerEndPoint,omitempty"`
	Extensions            *[]string `json:"extensions,omitempty" bson:"extensions,omitempty"`
	FriendlyName          string    `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
	Handle                float64   `json:"handle,omitempty" bson:"handle,omitempty"`
	HasPrivateKey         bool      `json:"hasPrivateKey,omitempty" bson:"hasPrivateKey,omitempty"`
	Issuer                string    `json:"issuer,omitempty" bson:"issuer,omitempty"`
	IssuingCertificateURL []string  `json:"issuingCertificateURL,omitempty" bson:"issuingCertificateURL,omitempty"`
	IssuerName            *struct {
		Name    string  `json:"name,omitempty" bson:"name,omitempty"`
		Oid     string  `json:"oid,omitempty" bson:"oid,omitempty"`
		RawData *string `json:"rawData,omitempty" bson:"rawData,omitempty"`
	} `json:"issuerName,omitempty" bson:"issuerName,omitempty"`
	NotAfter  time.Time `json:"notAfter,omitempty" bson:"notAfter,omitempty"`
	NotBefore time.Time `json:"notBefore,omitempty" bson:"notBefore,omitempty"`
	PolicyID  *string   `json:"policyId,omitempty" bson:"policyId,omitempty"`
	// PrivateKey *struct {
	// 	CspKeyContainerInfo  string  `json:"cspKeyContainerInfo,omitempty" bson:"cspKeyContainerInfo,omitempty"`
	// 	KeyExchangeAlgorithm string  `json:"keyExchangeAlgorithm,omitempty" bson:"keyExchangeAlgorithm,omitempty"`
	// 	KeySize              float64 `json:"keySize,omitempty" bson:"keySize,omitempty"`
	// 	LegalKeySizes        string  `json:"legalKeySizes,omitempty" bson:"legalKeySizes,omitempty"`
	// 	PersistKeyInCsp      bool    `json:"persistKeyInCsp,omitempty" bson:"persistKeyInCsp,omitempty"`
	// 	PublicOnly           bool    `json:"publicOnly,omitempty" bson:"publicOnly,omitempty"`
	// 	SignatureAlgorithm   string  `json:"signatureAlgorithm,omitempty" bson:"signatureAlgorithm,omitempty"`
	// } `json:"privateKey,omitempty" bson:"privateKey,omitempty"`
	PublicKey struct {
		EncodedKeyValue   string  `json:"encodedKeyValue,omitempty" bson:"encodedKeyValue,omitempty"`
		EncodedParameters string  `json:"encodedParameters,omitempty" bson:"encodedParameters,omitempty"`
		Key               *string `json:"key,omitempty" bson:"key,omitempty"`
		Oid               string  `json:"oid,omitempty" bson:"oid,omitempty"`
	} `json:"publicKey,omitempty" bson:"publicKey,omitempty"`
	RawData             []float64 `json:"rawData,omitempty" bson:"rawData,omitempty"`
	SendAsTrustedIssuer bool      `json:"sendAsTrustedIssuer,omitempty" bson:"sendAsTrustedIssuer,omitempty"`
	SerialNumber        string    `json:"serialNumber,omitempty" bson:"serialNumber,omitempty"`
	SignatureAlgorithm  struct {
		FriendlyName string `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
		Value        string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"signatureAlgorithm,omitempty" bson:"signatureAlgorithm,omitempty"`
	Subject     string `json:"subject,omitempty" bson:"subject,omitempty"`
	SubjectName struct {
		Name    string  `json:"name,omitempty" bson:"name,omitempty"`
		Oid     string  `json:"oid,omitempty" bson:"oid,omitempty"`
		RawData *string `json:"rawData,omitempty" bson:"rawData,omitempty"`
	} `json:"subjectName,omitempty" bson:"subjectName,omitempty"`
	Thumbprint string  `json:"thumbprint,omitempty" bson:"thumbprint,omitempty"`
	Version    float64 `json:"version,omitempty" bson:"version,omitempty"`
	ParentPath *string `json:"parentPath,omitempty" bson:"parentPath,omitempty"`
}

// type ServerCertInfoServersPulledFrom map[string][]string

type ServerCertInfoServersPulledFrom struct {
	ServerName       string   `json:"serverName,omitempty" bson:"serverName,omitempty"`
	CertificatePaths []string `json:"certificatePaths,omitempty" bson:"certificatePaths,omitempty"`
}

//
//
