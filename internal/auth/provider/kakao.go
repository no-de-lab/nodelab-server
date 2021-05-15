package provider

import (
	"fmt"
	"net/http"
)

const baseURL = "https://kapi.kakao.com"

// LoginKakao login by using kakao oauth
func LoginKakao(accessToken string) (string, error) {

	type kakaoResponse struct {
		Id string `json:"id"`
	}

	req, err := http.NewRequest("GET", baseURL+"/v2/user/me", nil)

	if err != nil {
		return "", fmt.Errorf("falied to make request of Kakao : %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := HTTPClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("falied to get response of Kakao : %w", err)
	}

	kr := new(kakaoResponse)
	err = GetJSON(resp, kr)

	if err != nil {
		return "", err
	}

	return kr.Id, nil
}
