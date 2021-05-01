package util

import (
	"encoding/json"
	"net/http"
	"time"
)

var HttpClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(r *http.Response, target interface{}) error {

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
