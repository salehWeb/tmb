package errors

import (
	"encoding/json"
	"fmt"
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
}


func NotFound(massage string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(Massage{Massage: massage})
	w.Write(bytes)
}


func MethodNotAllowed(w http.ResponseWriter, method string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(Massage{Massage: fmt.Sprintf("This %s Method Is Not Allowed", method)})
	w.Write(bytes)
}
