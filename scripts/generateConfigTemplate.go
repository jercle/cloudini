// fakeDataGen
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/jercle/cloudini/lib"
	"github.com/brianvoe/gofakeit/v7"
)

// gofakeit.Name()             // Markus Moen
// gofakeit.Email()            // alaynawuckert@kozey.biz
// gofakeit.Phone()            // (570)245-7485
// gofakeit.BS()               // front-end
// gofakeit.BeerName()         // Duvel
// gofakeit.Color()            // MediumOrchid
// gofakeit.Company()          // Moen, Pagac and Wuckert
// gofakeit.CreditCardNumber() // 4287271570245748
// gofakeit.HackerPhrase()     // Connecting the array won't do anything, we need to generate the haptic COM driver!
// gofakeit.JobTitle()         // Director
// gofakeit.CurrencyShort()    // USD

func main() {
	var (
		cldConfig lib.CldConfigRoot
	)
	envs := []string{"ENV1", "ENV2"}

	err := gofakeit.Struct(&cldConfig)
	lib.CheckFatalError(err)
	// jsonStr, _ := json.MarshalIndent(cldConfig, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(cldConfig.Cloudini.EncryptConfig)
	// cldConfig.Cloudini.EncryptConfig = false
	// fmt.Println(gofakeit.Username())
	// os.Exit(0)
	SophosConfig := generateSophosConfig(envs)
	Azure := generateAzureConfig(envs)
	CitrixCloud := generateCitrixCloudConfig(envs)
	ProxyConfig := generateProxyConfig()
	Domains := generateDomains(envs)
	cldConfig.SophosConfig = &SophosConfig
	cldConfig.Azure = &Azure
	cldConfig.CitrixCloud = &CitrixCloud
	cldConfig.ProxyConfig = &ProxyConfig
	cldConfig.Domains = &Domains

	jsonStr, _ := json.MarshalIndent(cldConfig, "", "  ")
	fmt.Println(string(jsonStr))

	// cldConfig.Azure.MultiTenantAuth.Tenants["T1"].CostExportsLocation =

	// id := gofakeit.UUID()
	// bs := gofakeit.AppName()
	// fmt.Println(bs)

}

func generateDomains(envs []string) map[string]string {
	fqdns := make(map[string]string)
	for _, env := range envs {
		fqdns[env] = gofakeit.DomainName()
	}
	return fqdns
}

func generateProxyConfig() map[string]lib.ProxyConfig {
	var proxyConfig map[string]lib.ProxyConfig
	jsonStr := `{
    "default": {
      "server": "proxy-lb.fq.d.n",
      "port": "8080",
      "enabled": true,
      "overrides": "localhost;127.0.0.;10.;192.168.*;*.azure.net;*.azure.com;*.windows.net;*.visualstudio.com;*.microsoft.com"
    },
    "other": {
      "server": "127.0.0.1",
      "port": "3128",
      "enabled": true
    }
  }`

	byteData := []byte(jsonStr)
	err := json.Unmarshal(byteData, &proxyConfig)
	lib.CheckFatalError(err)

	return proxyConfig
}

func generateCitrixCloudConfig(envs []string) lib.CitrixCloud {
	var ctxCldConf lib.CitrixCloud
	ctxCldConf.Environments = make(map[string]lib.CitrixCloudAccountConfig)

	for _, env := range envs {
		var ctxCldAcctConf lib.CitrixCloudAccountConfig
		err := gofakeit.Struct(&ctxCldAcctConf)
		lib.CheckFatalError(err)

		ctxCldConf.Environments[env] = ctxCldAcctConf
	}

	return ctxCldConf
}

func generateAzureConfig(envs []string) lib.AzureConfig {
	var azureConfig lib.AzureConfig
	tenants := make(map[string]lib.CldConfigTenantAuth)
	// rand.Seed(time.Now().UnixNano())
	randomEnv := rand.Intn(len(envs))

	for i, env := range envs {
		var tenantEnv lib.CldConfigTenantAuth
		err := gofakeit.Struct(&tenantEnv)
		lib.CheckFatalError(err)
		if i == randomEnv {
			tenantEnv.Default = true
		} else {
			tenantEnv.Default = false
		}
		tenantEnv.CostExportsLocation = "https://" + strings.ToLower(env) + "strgacct.blob.core.windows.net/cost-exports"
		tenantEnv.TenantName = env
		tenants[env] = tenantEnv
	}
	azureConfig.MultiTenantAuth.Tenants = tenants
	return azureConfig
}

func generateSophosConfig(envs []string) lib.SophosConfig {
	var sophosConfig lib.SophosConfig
	sophosConfig.Environments = make(map[string]lib.SophosEnvironment)
	for _, env := range envs {
		var sophosEnv lib.SophosEnvironment
		err := gofakeit.Struct(&sophosEnv)
		lib.CheckFatalError(err)
		sophosEnv.Hosts = []string{"sophos-host-1", "sophos-host-2"}
		sophosConfig.Environments[env] = sophosEnv
	}
	return sophosConfig
}
