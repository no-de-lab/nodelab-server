package util

import (
	"google.golang.org/api/oauth2/v2"
)

func LoginGoogle(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(HttpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
