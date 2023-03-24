package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Authentication(w http.ResponseWriter, r *http.Request) {

	var rPayload requestPayload
	err := app.readJSON(w, r, &rPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.models.User.GetByEmail(rPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("email not exist"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(rPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("password mismatch"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Err:  false,
		Msg:  fmt.Sprintf("Welcome %s", user.Email),
		Data: user,
	}

	err = app.writeJSON(w, http.StatusOK, payload, nil)
}
