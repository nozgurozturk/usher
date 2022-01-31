package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON is a helper to write JSON response
func JSON(w http.ResponseWriter, status int, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	_, _ = w.Write(buf.Bytes())
}
