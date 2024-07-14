package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

func ServeJSON(data interface{}) []byte {
	err, _ := json.Marshal(data)
	if err != nil {
		return err
	}
	return nil
}

func HttpResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
