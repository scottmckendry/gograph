package auth

import (
	"context"
	"goGraph/cache"
	"net/http"
	"os"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/joho/godotenv"
)

var Token string

func GetToken() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")

	var tokenCache cache.TokenCache
	tokenCache.File = "cache.json"

	client, err := public.New(clientId, public.WithCache(&tokenCache), public.WithAuthority("https://login.windows.net/"+tenantId))
	if err != nil {
		return err
	}

	ctx := context.Background()

	accounts, err := client.Accounts(ctx)
	var result public.AuthResult
	scopes := []string{"https://graph.microsoft.com/.default"}
	if err != nil {
		return err
	}
	if len(accounts) > 0 {
		// There may be more accounts; here we assume the first one is wanted.
		// AcquireTokenSilent returns a non-nil error when it can't provide a token.
		result, err = client.AcquireTokenSilent(ctx, scopes, public.WithSilentAccount(accounts[0]))
	}
	if err != nil || len(accounts) == 0 {
		// Failed cache, authenticate a user with another AcquireToken* method
		result, err = client.AcquireTokenInteractive(ctx, scopes)
		if err != nil {
			return err
		}
	}

	Token = result.AccessToken
	return nil
}

func MakeRequest(method string, url string, body []byte) (*http.Response, error) {
	if Token == "" {
		err := GetToken()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	return client.Do(req)
}
