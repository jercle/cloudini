package mongodb

import (
	"math/big"
	"net"
	"net/url"
	"time"

	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAllAzureResourcesAndVcpuCountsOptionsOptions struct {
	AzureResourcesDatabase      string
	CitrixDatabase              string
	AzResImageGalleryImagesColl string
	CitrixMachineCatalogsColl   string
}

type UpdateAllAzureResourcesAndVcpuCountsOptions struct {
	SkuListSubscription         string
	SkuListAuth                 lib.CldConfigTenantAuth
	Location                    string
	CostDataMonth               string
	CostDataBlobPrefix          string
	AzResSKUColl                *mongo.Collection
	AzResVcpuCountsColl         *mongo.Collection
	AzStorageAcctMinTlsVersions *mongo.Collection
	EnvOptCostingTenantsColl    *mongo.Collection
	EnvOptCostingSubsColl       *mongo.Collection
	EnvOptCostingResGrpsColl    *mongo.Collection
	EnvOptCostingResourcesColl  *mongo.Collection
	EnvOptCostingMetersColl     *mongo.Collection
	AzResTenantsColl            *mongo.Collection
	AzResResourceListColl       *mongo.Collection
	AzResGrpsListColl           *mongo.Collection
}

type UpdateEntraItemsOptions struct {
	EntraAppRegColl              *mongo.Collection
	EntraAppRegCredsExpiringColl *mongo.Collection
}

type UpdateEntraPimItemsOptions struct {
	EntraRoleEligibilityScheduleInstancesColl *mongo.Collection
	EntraRoleAssignmentScheduleInstancesColl  *mongo.Collection
}

type CertManagementConfig struct {
	ID                    string         `json:"_id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	PathsToIgnore         []string       `json:"pathsToIgnore,omitempty,omitzero" bson:"pathsToIgnore,omitempty,omitzero"`
	FriendlyNamesToIgnore []FriendlyName `json:"friendlyNamesToIgnore,omitempty,omitzero" bson:"friendlyNamesToIgnore,omitempty,omitzero"`
	SubjectNamesToIgnore  []SubjectName  `json:"subjectNamesToIgnore,omitempty,omitzero" bson:"subjectNamesToIgnore,omitempty,omitzero"`
	IssuersToIgnore       []Issuer       `json:"issuersToIgnore,omitempty,omitzero" bson:"issuersToIgnore,omitempty,omitzero"`
	UrlsToWatch           []struct {
		URL    string `json:"url,omitempty,omitzero" bson:"url,omitempty,omitzero"`
		Name   string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Tenant string `json:"tenant,omitempty,omitzero" bson:"tenant,omitempty,omitzero"`
	} `json:"urlsToWatch,omitempty,omitzero" bson:"urlsToWatch,omitempty,omitzero"`
}

type Issuer struct {
	Issuer               string `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
	IsRegex              bool   `json:"isRegex,omitempty,omitzero" bson:"isRegex,omitempty,omitzero"`
	RegExCaseInsensitive bool   `json:"regExCaseInsensitive,omitempty,omitzero" bson:"regExCaseInsensitive,omitempty,omitzero"`
	RegExStartsWith      bool   `json:"regExStartsWith,omitempty,omitzero" bson:"regExStartsWith,omitempty,omitzero"`
	RegExEndsWith        bool   `json:"regExEndsWith,omitempty,omitzero" bson:"regExEndsWith,omitempty,omitzero"`
}

type SubjectName struct {
	SubjectName          string `json:"subjectName,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
	IsRegex              bool   `json:"isRegex,omitempty,omitzero" bson:"isRegex,omitempty,omitzero"`
	RegExCaseInsensitive bool   `json:"regExCaseInsensitive,omitempty,omitzero" bson:"regExCaseInsensitive,omitempty,omitzero"`
	RegExStartsWith      bool   `json:"regExStartsWith,omitempty,omitzero" bson:"regExStartsWith,omitempty,omitzero"`
	RegExEndsWith        bool   `json:"regExEndsWith,omitempty,omitzero" bson:"regExEndsWith,omitempty,omitzero"`
}

type FriendlyName struct {
	FriendlyName         string `json:"friendlyName,omitempty,omitzero" bson:"friendlyName,omitempty,omitzero"`
	IsRegex              bool   `json:"isRegex,omitempty,omitzero" bson:"isRegex,omitempty,omitzero"`
	RegExCaseInsensitive bool   `json:"regExCaseInsensitive,omitempty,omitzero" bson:"regExCaseInsensitive,omitempty,omitzero"`
	RegExStartsWith      bool   `json:"regExStartsWith,omitempty,omitzero" bson:"regExStartsWith,omitempty,omitzero"`
	RegExEndsWith        bool   `json:"regExEndsWith,omitempty,omitzero" bson:"regExEndsWith,omitempty,omitzero"`
}

type WebsiteCertificateMinimal struct {
	SerialNumber          *big.Int   `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	Issuer                string     `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
	Subject               string     `json:"subject,omitempty,omitzero" bson:"subject,omitempty,omitzero"`
	NotBefore             time.Time  `json:"notBefore,omitempty,omitzero" bson:"notBefore,omitempty,omitzero"`
	NotAfter              time.Time  `json:"notAfter,omitempty,omitzero" bson:"notAfter,omitempty,omitzero"`
	BasicConstraintsValid bool       `json:"basicConstraintsValid,omitempty,omitzero" bson:"basicConstraintsValid,omitempty,omitzero"`
	IsCA                  bool       `json:"isCA,omitempty,omitzero" bson:"isCA,omitempty,omitzero"`
	IssuingCertificateURL []string   `json:"issuingCertificateURL,omitempty,omitzero" bson:"issuingCertificateURL,omitempty,omitzero"`
	DNSNames              []string   `json:"dnsNames,omitempty,omitzero" bson:"dnsNames,omitempty,omitzero"`
	EmailAddresses        []string   `json:"emailAddresses,omitempty,omitzero" bson:"emailAddresses,omitempty,omitzero"`
	IPAddresses           []net.IP   `json:"ipAddresses,omitempty,omitzero" bson:"ipAddresses,omitempty,omitzero"`
	URIs                  []*url.URL `json:"uris,omitempty,omitzero" bson:"uris,omitempty,omitzero"`
	CRLDistributionPoints []string   `json:"crlDistributionPoints,omitempty,omitzero" bson:"crlDistributionPoints,omitempty,omitzero"`
}

//
//

type AWSIngestCounts struct {
	Counts           []AWSIngestCount `json:"counts" bson:"counts"`
	TotalLogs        float64          `json:"totalLogs" bson:"totalLogs"`
	TotalSQSMessages float64          `json:"totalSQSMessages" bson:"totalSQSMessages"`
	Environment      string           `json:"environment" bson:"environment"`
	Monitor          string           `json:"monitor" bson:"monitor"`
	ID               string           `json:"id" bson:"_id"`
	LastDBSync       time.Time        `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty"`
}

//
//

type AWSIngestCount struct {
	LogType               string  `json:"logType" bson:"logType"`
	Count                 float64 `json:"count" bson:"count"`
	PercentageOfTotalLogs float64 `json:"percentageOfTotalLogs" bson:"percentageOfTotalLogs"`
}
