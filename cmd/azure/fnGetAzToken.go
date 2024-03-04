package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/charmbracelet/log"
)

type Request struct {
	Url     string
	Outfile string
}

type TokenData struct {
	Token     string
	ExpiresOn string
}

func GetServicePrincipalToken() TokenData {
	ctx := context.Background()

	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
		EnableCAE: true,
	}

	envCred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	tokenResponse, err := envCred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	token := TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return token
}

func GetAzCliToken() TokenData {
	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
		EnableCAE: true,
	}

	cliCred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	tokenResponse, err := cliCred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	token := TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return token
}
