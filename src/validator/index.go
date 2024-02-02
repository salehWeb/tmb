package validator

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/salehWeb/chat-app/server/src/errors"
)

func IsNotEmail(email string, w http.ResponseWriter) bool {

	if len(email) < 1 {
		errors.BadRequest("No Email Found", w)
		return true
	}

	_, err := mail.ParseAddress(email)

	if err != nil {
		errors.BadRequest("UnValid Email Address", w)
		return true
	}

	return false
}

func IsNotLen(short int, long int, v string, name string, w http.ResponseWriter) bool {
	length := len(v)

	if length < 1 {
		errors.BadRequest(fmt.Sprintf("No %s Found", name), w)
		return true
	}

	if length >= long {
		errors.BadRequest(fmt.Sprintf("%s Is To Long", name), w)
		return true
	}

	if length <= short {
		errors.BadRequest(fmt.Sprintf("%s Is To Short", name), w)
		return true
	}

	return false
}
