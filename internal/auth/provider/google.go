package provider

import (
	"context"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// LoginGoogle get IDToken of Google
func LoginGoogle(idToken string) (*oauth2.Tokeninfo, error) {
	ctx := context.Background()
	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(HTTPClient))
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
