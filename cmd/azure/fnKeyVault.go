package azure

import (
	"encoding/json/v2"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jercle/cloudini/lib"
	"github.com/r3labs/diff/v3"
)

func GetAllKVSecretsInTenant(token *lib.AzureMultiAuthToken, kvToken *lib.AzureMultiAuthToken) (allSecrets []KeyVaultSecretStored) {
	subs, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)

	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	for _, sub := range subs {
		wg.Go(func() {
			kvs := ListKeyVaultsForSub(sub, token)

			if len(kvs) == 0 {
				return
			}

			for _, kv := range kvs {
				if strings.HasPrefix(kv.Name, "pkr") {
					continue
				}
				secrets := KeyVaultListSecrets(kv, kvToken)
				mut.Lock()
				allSecrets = append(allSecrets, secrets...)
				mut.Unlock()
			}
		})
	}

	wg.Wait()

	return
}

//
//

func ListKeyVaultsForSub(sub lib.FetchedSubscription, token *lib.AzureMultiAuthToken) (kvs []KeyVault) {
	fmt.Println("Fetching KVs for", sub.DisplayName)
	urlString := "https://management.azure.com/subscriptions/" + sub.SubscriptionID + "/providers/Microsoft.KeyVault/vaults?api-version=2024-11-01"
	res, err := HttpGetErrLogToCache(urlString, *token)
	lib.CheckFatalError(err)

	var resData ListKeyVaultsForSubResponse
	json.Unmarshal(res, &resData)

	// interface{}

	for _, kv := range resData.Value {
		curr := kv
		curr.SubscriptionName = sub.DisplayName
		curr.SubscriptionId = sub.SubscriptionID
		curr.TenantName = token.TenantName
		curr.TenantId = token.TenantId

		kvs = append(kvs, curr)
	}

	for resData.NextLink != "" {
		// fmt.Println(resData.NextLink)

		res, err := HttpGet(resData.NextLink, *token)
		lib.CheckFatalError(err)
		resData = ListKeyVaultsForSubResponse{}
		json.Unmarshal(res, &resData)

		// kvs = append(kvs, resData.Value...)
		for _, kv := range resData.Value {
			curr := kv
			curr.SubscriptionName = sub.DisplayName
			curr.SubscriptionId = sub.SubscriptionID
			curr.TenantName = token.TenantName
			curr.TenantId = token.TenantId

			kvs = append(kvs, curr)
		}
	}

	return
}

//
//

func ListKeyVaultsForTenant(token *lib.AzureMultiAuthToken) (kvs []KeyVault) {
	subs, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)

	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	fmt.Println("Fetching KVs for", token.TenantName)

	for _, sub := range subs {
		wg.Go(func() {
			subVaults := ListKeyVaultsForSub(sub, token)
			mut.Lock()
			kvs = append(kvs, subVaults...)
			mut.Unlock()
		})
	}
	wg.Wait()

	return
}

//
//

func GetKeyVault(kvId string, token *lib.AzureMultiAuthToken) (kv KeyVault) {
	urlString := "https://management.azure.com" + kvId + "?api-version=2024-11-01"
	res, err := HttpGetErrLogToCache(urlString, *token)
	lib.CheckFatalError(err)
	// fmt.Println(string(res))

	err = json.Unmarshal(res, &kv)

	// lib.JsonMarshalAndPrint(kv)

	return
}

//
//

func KeyVaultListSecrets(keyVault KeyVault, token *lib.AzureMultiAuthToken) (processedSecrets []KeyVaultSecretStored) {
	apiVersion := "?api-version=2025-07-01"
	// apiVersion := "?api-version=7.4"
	urlString := keyVault.Properties.VaultURI + "secrets" + apiVersion

	fmt.Println("Listing secrets for", keyVault.Name)

	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var secrets []KeyVaultSecret

	var resData ListKeyVaultSecretsResponse
	json.Unmarshal(res, &resData)
	secrets = append(secrets, resData.Value...)

	for resData.NextLink != "" {
		res, err := HttpGet(resData.NextLink, *token)
		lib.CheckFatalError(err)
		resData = ListKeyVaultSecretsResponse{}
		json.Unmarshal(res, &resData)
		secrets = append(secrets, resData.Value...)
	}

	keyVaultUrl := keyVault.Properties.VaultURI
	for _, s := range secrets {
		if s.ContentType == "application/x-pkcs12" || s.ContentType == "application/x-pem-file" {
			continue
		}

		curr := KeyVaultSecretStored{
			Name:             strings.Replace(s.ID, keyVaultUrl+"secrets/", "", 1),
			Id:               s.ID,
			KeyVaultId:       keyVault.ID,
			KeyVaultUrl:      keyVault.Properties.VaultURI,
			KeyVaultName:     keyVault.Name,
			TenantName:       keyVault.TenantName,
			SubscriptionName: keyVault.SubscriptionName,
			Type:             "Secret",
			ContentType:      s.ContentType,
		}

		if s.Attributes.Exp != nil {
			curr.Expiration = time.Unix(*s.Attributes.Exp, 0)
		}

		processedSecrets = append(processedSecrets, curr)
	}

	certs := KeyVaultListCerts(keyVault, token)

	for _, c := range certs {
		curr := KeyVaultSecretStored{
			Name:             strings.Replace(c.ID, keyVaultUrl+"certificates/", "", 1),
			Id:               c.ID,
			KeyVaultId:       keyVault.ID,
			KeyVaultUrl:      keyVault.Properties.VaultURI,
			KeyVaultName:     keyVault.Name,
			TenantName:       keyVault.TenantName,
			SubscriptionName: keyVault.SubscriptionName,
			Type:             "Certificate",
			ContentType:      "Certificate",
			Expiration:       time.Unix(c.Attributes.Exp, 0),
		}
		processedSecrets = append(processedSecrets, curr)
	}
	return
}

//
//

func KeyVaultListCerts(keyVault KeyVault, token *lib.AzureMultiAuthToken) (certs []KeyVaultCertificateMin) {
	apiVersion := "?api-version=2025-07-01"
	// apiVersion := "?api-version=7.4"
	urlString := keyVault.Properties.VaultURI + "certificates" + apiVersion

	// fmt.Println("Listing secrets for", keyVault.Name)

	res, err := HttpGetErrLogToCache(urlString, *token)
	lib.CheckFatalError(err)

	// var certs []KeyVaultCertificate

	var resData ListKeyVaultCertsResponse
	json.Unmarshal(res, &resData)
	certs = append(certs, resData.Value...)

	for resData.NextLink != "" {
		res, err := HttpGet(resData.NextLink, *token)
		lib.CheckFatalError(err)
		resData = ListKeyVaultCertsResponse{}
		json.Unmarshal(res, &resData)
		certs = append(certs, resData.Value...)
	}

	return
}

//
//

func KeyVaultGetCert(certId string, token *lib.AzureMultiAuthToken) {
	// func KeyVaultGetCert(keyVault KeyVault, token *lib.AzureMultiAuthToken) KeyVaultSecretStored {
	apiVersion := "?api-version=2025-07-01"
	// apiVersion := "?api-version=7.4"
	urlString := certId + apiVersion

	// fmt.Println("Listing secrets for", keyVault.Name)

	res, err := HttpGetErrLogToCache(urlString, *token)
	lib.CheckFatalError(err)

	fmt.Println(string(res))

	// var certs []KeyVaultCertificate

	// var resData ListKeyVaultCertsResponse
	// json.Unmarshal(res, &resData)
	// certs = append(certs, resData.Value...)

	// for resData.NextLink != "" {
	// 	res, err := HttpGet(resData.NextLink, *token)
	// 	lib.CheckFatalError(err)
	// 	resData = ListKeyVaultCertsResponse{}
	// 	json.Unmarshal(res, &resData)
	// 	certs = append(certs, resData.Value...)
	// }

	return
}

//
//

func KeyVaultListSecretsAndCertsForAllConfiguredTenants() (allSecrets []KeyVaultSecretStored) {
	cldConfig := lib.GetCldConfig(nil)
	azConfigs := cldConfig.Azure.MultiTenantAuth.Tenants

	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	for tName, tConfig := range azConfigs {
		if tConfig.IsB2C {
			continue
		}

		wg.Go(func() {
			token, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
				// GetWriteToken: true,
			}, nil)
			lib.CheckFatalError(err)

			kvToken, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				Scope:      "keyvault",
				TenantName: tName,
				// GetWriteToken: true,
			}, nil)
			lib.CheckFatalError(err)

			secrets := GetAllKVSecretsInTenant(token, kvToken)
			mut.Lock()
			allSecrets = append(allSecrets, secrets...)
			mut.Unlock()

		})

	}

	wg.Wait()

	return
}

//
//

func KeyVaultAddSubnetForAllInConfiguredTenants(subnetId string, configuredTenantNames []string) (comparisons []KeyVaultUpdateComparison) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)

	for _, token := range tokenReq {
		if len(configuredTenantNames) > 0 {
			if !slices.Contains(configuredTenantNames, token.TenantName) {
				continue
			}
		}
		// fmt.Println(token.TenantName)
		// continue
		wg.Go(func() {
			kvs := ListKeyVaultsForTenant(&token)
			var wgkvs sync.WaitGroup
			for _, k := range kvs {
				wgkvs.Go(func() {

					var (
						update    KeyVault
						updatedKv KeyVault
					)

					kv := GetKeyVault(k.ID, &token)

					urlString := "https://management.azure.com" + kv.ID + "?api-version=2024-11-01"

					newSubnetRule := KeyVaultNetworkACLVirtualNetworkRule{
						ID: subnetId,
						// IgnoreMissingVnetServiceEndpoint: true,
					}

					vnetRules := kv.Properties.NetworkAcls.VirtualNetworkRules
					vnetRules = append(vnetRules, newSubnetRule)
					update.Properties.NetworkAcls.VirtualNetworkRules = removeDuplicateVnetRules(vnetRules)

					// accessPolicies := kv.Properties.AccessPolicies
					// newAccessPol := KeyVaultAccessPolicy{
					// 	TenantID: token.TenantId,
					// 	ObjectID: "",
					// 	Permissions: KeyVaultAccessPolicyPermissions{
					// 		Certificates: []string{"get", "list"},
					// 		Keys:         []string{"get", "list"},
					// 		Secrets:      []string{"get", "list"},
					// 	},
					// }
					// accessPolicies = append(accessPolicies, newAccessPol)
					// update.Properties.AccessPolicies = removeDuplicateAccessPolicies(accessPolicies)

					bodyStr, err := json.Marshal(update)
					resBody, _, err := HttpPatchErrLogToCache(urlString, string(bodyStr), token)
					// lib.CheckFatalError(err)
					if err != nil {
						fmt.Println(err)
					}

					json.Unmarshal(resBody, &updatedKv)

					d, err := diff.NewDiffer(diff.Filter(func(path []string, parent reflect.Type, field reflect.StructField) bool {
						return field.Name != "SystemData"
					}))

					cl, err := d.Diff(kv, updatedKv)
					lib.CheckFatalError(err)

					c := KeyVault{}
					patchlog := diff.Patch(cl, &c)

					comparison := KeyVaultUpdateComparison{
						Id:         kv.ID,
						Original:   kv,
						Updated:    updatedKv,
						Diff:       patchlog,
						DiffString: cmp.Diff(kv, updatedKv),
					}

					mut.Lock()
					comparisons = append(comparisons, comparison)
					mut.Unlock()
				})
			}
			wgkvs.Wait()
		})
	}

	wg.Wait()

	return
}

//
//

func removeDuplicateVnetRules(subnets []KeyVaultNetworkACLVirtualNetworkRule) []KeyVaultNetworkACLVirtualNetworkRule {
	seen := make(map[string]bool)
	unique := []KeyVaultNetworkACLVirtualNetworkRule{}

	for _, s := range subnets {
		if !seen[strings.ToLower(s.ID)] {
			seen[strings.ToLower(s.ID)] = true
			unique = append(unique, s)
		}
	}
	return unique
}

func removeDuplicateAccessPolicies(accessPolicies []KeyVaultAccessPolicy) []KeyVaultAccessPolicy {
	seen := make(map[string]bool)
	unique := []KeyVaultAccessPolicy{}

	for _, a := range accessPolicies {
		if !seen[strings.ToLower(a.ObjectID)] {
			seen[strings.ToLower(a.ObjectID)] = true
			unique = append(unique, a)
		}
	}
	return unique
}
