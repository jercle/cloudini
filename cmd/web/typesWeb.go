package web

import (
	"math/big"
	"net"
	"net/url"
	"time"

	"github.com/jercle/cloudini/lib"
)

type WebsiteCertificateMinimal struct {
	SerialNumber          *big.Int                              `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	Issuer                string                                `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
	Subject               string                                `json:"subject,omitempty,omitzero" bson:"subject,omitempty,omitzero"`
	NotBefore             time.Time                             `json:"notBefore,omitempty,omitzero" bson:"notBefore,omitempty,omitzero"`
	NotAfter              time.Time                             `json:"notAfter,omitempty,omitzero" bson:"notAfter,omitempty,omitzero"`
	BasicConstraintsValid bool                                  `json:"basicConstraintsValid,omitempty,omitzero" bson:"basicConstraintsValid,omitempty,omitzero"`
	IsCA                  bool                                  `json:"isCA,omitempty,omitzero" bson:"isCA,omitempty,omitzero"`
	IssuingCertificateURL []string                              `json:"issuingCertificateURL,omitempty,omitzero" bson:"issuingCertificateURL,omitempty,omitzero"`
	DNSNames              []string                              `json:"dnsNames,omitempty,omitzero" bson:"dnsNameList,omitempty,omitzero"`
	EmailAddresses        []string                              `json:"emailAddresses,omitempty,omitzero" bson:"emailAddresses,omitempty,omitzero"`
	IPAddresses           []net.IP                              `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	URIs                  []*url.URL                            `json:"uris,omitempty,omitzero" bson:"uris,omitempty,omitzero"`
	CRLDistributionPoints []string                              `json:"crlDistributionPoints,omitempty,omitzero" bson:"crlDistributionPoints,omitempty,omitzero"`
	LastDBSync            *time.Time                            `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
	TenantName            string                                `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	FriendlyName          string                                `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
	ServersPulledFrom     []lib.ServerCertInfoServersPulledFrom `json:"serversPulledFrom,omitempty" bson:"serversPulledFrom,omitempty"`
}

// type ServerCertInfo struct {
// 	ID                       string                            `json:"id,omitempty" bson:"_id,omitempty"`
// 	IgnoreExpiration         bool                              `json:"ignoreExpiration,omitempty"  bson:"ignoreExpiration"`
// 	LastDBSync               *time.Time                        `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
// 	LastServerSync           time.Time                         `json:"lastServerSync,omitempty"  bson:"lastServerSync,omitempty"`
// 	PulledFromServer         *string                           `json:"pulledFromServer,omitempty" bson:"pulledFromServer,omitempty"`
// 	ServersPulledFrom        []ServerCertInfoServersPulledFrom `json:"serversPulledFrom,omitempty" bson:"serversPulledFrom,omitempty"`
// 	TenantName               string                            `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
// 	RelatedCertAuthData      *CertAuthorityCertInfo            `json:"relatedCertAuthData,omitempty" bson:"relatedCertAuthData,omitempty"`
// 	Archived                 bool                              `json:"archived,omitempty" bson:"archived,omitempty"`
// 	DnsNameList              *[]string                         `json:"dnsNameList,omitempty" bson:"dnsNameList,omitempty"`
// 	EnhancedKeyUsageList     *[]string                         `json:"enhancedKeyUsageList,omitempty" bson:"enhancedKeyUsageList,omitempty"`
// 	EnrollmentPolicyEndPoint *struct {
// 		AuthenticationType float64 `json:"authenticationType,omitempty" bson:"authenticationType,omitempty"`
// 		URL                *string `json:"url,omitempty" bson:"url,omitempty"`
// 	} `json:"EnrollmentPolicyEndPoint,omitempty" bson:"EnrollmentPolicyEndPoint,omitempty"`
// 	EnrollmentServerEndPoint *struct {
// 		AuthenticationType float64 `json:"authenticationType,omitempty" bson:"authenticationType,omitempty"`
// 		URL                *string `json:"url,omitempty" bson:"url,omitempty"`
// 	} `json:"enrollmentServerEndPoint,omitempty" bson:"enrollmentServerEndPoint,omitempty"`
// 	Extensions    *[]string `json:"extensions,omitempty" bson:"extensions,omitempty"`
// 	FriendlyName  string    `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
// 	Handle        float64   `json:"handle,omitempty" bson:"handle,omitempty"`
// 	HasPrivateKey bool      `json:"hasPrivateKey,omitempty" bson:"hasPrivateKey,omitempty"`
// 	Issuer        string    `json:"issuer,omitempty" bson:"issuer,omitempty"`
// 	IssuerName    *struct {
// 		Name    string  `json:"name,omitempty" bson:"name,omitempty"`
// 		Oid     string  `json:"oid,omitempty" bson:"oid,omitempty"`
// 		RawData *string `json:"rawData,omitempty" bson:"rawData,omitempty"`
// 	} `json:"issuerName,omitempty" bson:"issuerName,omitempty"`
// 	NotAfter  time.Time `json:"notAfter,omitempty" bson:"notAfter,omitempty"`
// 	NotBefore time.Time `json:"notBefore,omitempty" bson:"notBefore,omitempty"`
// 	PolicyID  *string   `json:"policyId,omitempty" bson:"policyId,omitempty"`
// 	// PrivateKey *struct {
// 	// 	CspKeyContainerInfo  string  `json:"cspKeyContainerInfo,omitempty" bson:"cspKeyContainerInfo,omitempty"`
// 	// 	KeyExchangeAlgorithm string  `json:"keyExchangeAlgorithm,omitempty" bson:"keyExchangeAlgorithm,omitempty"`
// 	// 	KeySize              float64 `json:"keySize,omitempty" bson:"keySize,omitempty"`
// 	// 	LegalKeySizes        string  `json:"legalKeySizes,omitempty" bson:"legalKeySizes,omitempty"`
// 	// 	PersistKeyInCsp      bool    `json:"persistKeyInCsp,omitempty" bson:"persistKeyInCsp,omitempty"`
// 	// 	PublicOnly           bool    `json:"publicOnly,omitempty" bson:"publicOnly,omitempty"`
// 	// 	SignatureAlgorithm   string  `json:"signatureAlgorithm,omitempty" bson:"signatureAlgorithm,omitempty"`
// 	// } `json:"privateKey,omitempty" bson:"privateKey,omitempty"`
// 	PublicKey struct {
// 		EncodedKeyValue   string  `json:"encodedKeyValue,omitempty" bson:"encodedKeyValue,omitempty"`
// 		EncodedParameters string  `json:"encodedParameters,omitempty" bson:"encodedParameters,omitempty"`
// 		Key               *string `json:"key,omitempty" bson:"key,omitempty"`
// 		Oid               string  `json:"oid,omitempty" bson:"oid,omitempty"`
// 	} `json:"publicKey,omitempty" bson:"publicKey,omitempty"`
// 	RawData             []float64 `json:"rawData,omitempty" bson:"rawData,omitempty"`
// 	SendAsTrustedIssuer bool      `json:"sendAsTrustedIssuer,omitempty" bson:"sendAsTrustedIssuer,omitempty"`
// 	SerialNumber        string    `json:"serialNumber,omitempty" bson:"serialNumber,omitempty"`
// 	SignatureAlgorithm  struct {
// 		FriendlyName string `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
// 		Value        string `json:"value,omitempty" bson:"value,omitempty"`
// 	} `json:"signatureAlgorithm,omitempty" bson:"signatureAlgorithm,omitempty"`
// 	Subject     string `json:"subject,omitempty" bson:"subject,omitempty"`
// 	SubjectName struct {
// 		Name    string  `json:"name,omitempty" bson:"name,omitempty"`
// 		Oid     string  `json:"oid,omitempty" bson:"oid,omitempty"`
// 		RawData *string `json:"rawData,omitempty" bson:"rawData,omitempty"`
// 	} `json:"subjectName,omitempty" bson:"subjectName,omitempty"`
// 	Thumbprint string  `json:"thumbprint,omitempty" bson:"thumbprint,omitempty"`
// 	Version    float64 `json:"version,omitempty" bson:"version,omitempty"`
// 	ParentPath *string `json:"parentPath,omitempty" bson:"parentPath,omitempty"`
// }
