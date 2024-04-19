// Azure DevOps authentication

package main

import (
	"context"
	"fmt"

	"github.com/jercle/cloudini/lib"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

func main() {
	ctx := context.Background()
	clientId := ""
	tenantId := ""

	client, err := public.New(clientId, public.WithAuthority("https://login.microsoftonline.com/"+tenantId))
	lib.CheckFatalError(err)

	// fmt.Println(client)
	// accounts, err := client.Accounts(ctx)
	// fmt.Println(accounts)

	scopes := []string{
		"User.ReadAll",
	}

	// auth, err := client.AcquireTokenInteractive(ctx, scopes)
	auth, err := client.AcquireTokenByDeviceCode(ctx, scopes)
	lib.CheckFatalError(err)

	fmt.Println(auth)
}
