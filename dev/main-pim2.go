// Azure DevOps authentication

package main

import (
	"context"
	"fmt"

	"github.com/jercle/azg/lib"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

func main() {
	ctx := context.Background()
	clientId := "4ca69554-1c1c-4eb5-84d6-a7e61551929a"
	tenantId := "e9f4bce2-7308-461a-91ce-3f663d079f47"

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
