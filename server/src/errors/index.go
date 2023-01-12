package errors

import (
	"encoding/json"
	"net/http"
)

type Massage struct {
	Massage string `json:"massage"`
}

func BadRequest(massage string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(Massage{Massage: massage})
	w.Write(bytes)
	return
}
