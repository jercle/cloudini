package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/azuredns"
	"github.com/go-acme/lego/v4/registration"
	"github.com/jercle/azg/lib"
	// "github.com/jercle/lego/providers/dns/azuredns"
)

// You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func main() {

	// Create a user. New accounts need an email and private key to start.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	myUser := MyUser{
		Email: "you@yours.com",
		key:   privateKey,
	}

	config := lego.NewConfig(&myUser)

	// azToken, err := azure.GetTenantSPToken("REDDTQ", lib.MultiAuthTokenRequestOptions{})
	// _ = azToken
	// lib.CheckFatalError(err)

	// fmt.Println(azToken)
	// os.Exit(0)

	var conf azuredns.Config
	// var tok exported.TokenCredential

	// cldConf := lib.GetCldConfig(&lib.CldConfigOptions{})

	// tenantConfig := cldConf.Azure.MultiTenantAuth.Tenants["REDDTQ"]

	// fmt.Println(tenantConfig)

	// creds, err := azidentity.

	conf.ClientID = os.Getenv("AZURE_CLIENT_ID")
	conf.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	conf.ResourceGroup = os.Getenv("AZURE_RESOURCE_GROUP")
	conf.TenantID = os.Getenv("AZURE_TENANT_ID")
	conf.SubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")

	cred, err := azidentity.NewClientSecretCredential(conf.TenantID, conf.ClientID, conf.ClientSecret, nil)
	_ = cred
	// azDnsProvider, err := azuredns.NewDNSProviderPublic(&conf, cred)
	// lib.CheckFatalError(err)

	// // fmt.Println(azDnsProvider)
	// jsonBytes, _ := json.MarshalIndent(conf, "", "  ")
	// lib.PrintJsonBytes(jsonBytes)

	err = azDnsProvider.Present("nothing.acmetest.stkcat.dev", "", "")
	lib.CheckFatalError(err)

	joelTest("testing", "testing2", "testing3", "testing4")
	joelTest("testing", "testing2", "testing3")
	joelTest("testing", "testing2")

	os.Exit(0)

	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// We specify an HTTP port of 5002 and an TLS port of 5001 on all interfaces
	// because we aren't running as root and can't bind a listener to port 80 and 443
	// (used later when we attempt to pass challenges). Keep in mind that you still
	// need to proxy challenge traffic to port 5002 and 5001.
	// err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", "5001"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Challenge.SetDNS01Provider(dns01.)

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	myUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{"nothing.stkcat.dev"},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	fmt.Printf("%#v\n", certificates)

	// ... all done.

}

func joelTest(name string, day string, options ...string) {
	_ = name
	_ = day
	fmt.Println(options)
}
