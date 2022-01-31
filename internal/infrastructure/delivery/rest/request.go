package rest

import (
	"encoding/json"
	"net/http"
)

// ParseBody parses the body of the request
func ParseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
