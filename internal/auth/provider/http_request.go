package provider

import (
	"encoding/json"
	"net/http"
	"time"
)

// HTTPClient define client of http
var HTTPClient = &http.Client{Timeout: 10 * time.Second}

// GetJSON make request data to json
func GetJSON(r *http.Response, target interface{}) error {

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}
