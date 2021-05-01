package util

import (
	"fmt"
	"net/http"
)

// LoginKakao login by using kakao oauth
func LoginKakao(accessToken string) (string, error) {
	baseURL := "https://kapi.kakao.com"

	type _kakaoResponse struct {
		id string
	}

	req, err := http.NewRequest("GET", baseURL+"/v2/user/me", nil)

	if err != nil {
		return "", fmt.Errorf(("Falied to make request of Kakao"))
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := HttpClient.Do(req)

	if err != nil {
		return "", fmt.Errorf(("Falied to get response of Kakao"))
	}

	kakaoResponse := new(_kakaoResponse)
	GetJson(resp, kakaoResponse)

	return kakaoResponse.id, nil
}
