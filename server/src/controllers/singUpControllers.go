package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/salehWeb/chat-app/server/src/dto"
	"github.com/salehWeb/chat-app/server/src/errors"
	"github.com/salehWeb/chat-app/server/src/helpers"
	"github.com/salehWeb/chat-app/server/src/initializers"
	"github.com/salehWeb/chat-app/server/src/models"
)

func SingUpControllers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// check if There No Data
		if r.Body == nil {
			errors.BadRequest("No Data Found In Request", w)
			return
		}

		// check if data is formatted correctly

		var data dto.SingUp
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			errors.BadRequest("Invalid Type Of Request Data", w)
			return
		}

		// check if same Email found
		var isFound struct {
			ID *uint
		}

		initializers.DB.Raw("SELECT id FROM users WHERE email = ?", data.Email).Scan(&isFound)
		if isFound.ID != nil {
			errors.BadRequest(fmt.Sprintf("User With This Email \"%s\" is Already Exists Try Login", data.Email), w)
			return
		}

		// hash password
		hash := helpers.HashPassword(data.Password)

		// create account

		type accountData struct {
			ID        uint   `json:"id"`
			Email     string `json:"email"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		}

		var account accountData

		initializers.DB.Create(&models.User{
			Password:  hash,
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
		}).Select("ID", "Email", "FirstName", "LastName").Scan(&account)

		// initializers.DB.Raw("SELECT ID, Email, FirstName, LastName FROM User WHERE ID = ?", )
		// set token in cookie
		cookie := http.Cookie{
			Name:     "accessToken",
			Value:    helpers.NewToken(account.ID),
			MaxAge:   int(time.Now().Add(time.Duration(time.Now().Year())).UTC().Unix()),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}

		http.SetCookie(w, &cookie)
		// return id, firstName, lastName, email
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(account)
		w.Write(bytes)
		return
		// done
	}

	errors.BadRequest(fmt.Sprintf("UnSuppurated request Method \"%s\"", r.Method), w)
}
