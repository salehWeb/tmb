package controllers

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/src/err"
)

func SingUpControllers(w http.ResponseWriter, r *http.Request) {



	if r.Method == http.MethodPost {
		// check if There No Data
		if r.Body == nil {
			err.BadRequest("No Data Found In Request", w)
			return
		}

		// check if data is formatted correctly

		// check if same Email found

		// hash password

		// create account

		// set token in cookie

		// return id, firstName, lastName, email

		// done

		return
	}


	if r.Method == http.MethodGet {
		err.BadRequest("It Should Be true", w)
		return
	}
}
