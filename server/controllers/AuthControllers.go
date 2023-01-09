package controllers

import (
	"net/http"
)

func AuthControllers(w http.ResponseWriter, r *http.Request) {

	 if r.Method == http.MethodGet {
		 return
	 }

	 if r.Method == http.MethodPost {
		 return
	 }

	 if r.Method == http.MethodPut {
		 return
	 }

	 if r.Method == http.MethodDelete {
		 return
	 }

}
