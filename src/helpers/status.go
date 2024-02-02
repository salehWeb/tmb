package helpers

import (
	"encoding/json"
	"net/http"
)


type Res struct {
	Massage string `json:"massage"`
	Data interface{} `json:"data"`
}


func Ok(val interface{}, massage string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := Res{
		Massage: massage,
		Data: val,
	}

	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}

func Created(val interface{}, massage string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	res := Res{
		Massage: massage,
		Data: val,
	}

	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}
