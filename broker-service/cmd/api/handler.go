package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Err: false,
		Msg: "from broker",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		log.Print("error reading json ", err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)

	case "log":

	default:
		app.errorJSON(w, errors.New("invalid action"))
		return
	}

}

func (app *Config) log() {

}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// build the payload that will be sent to microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// send the payload to the microservice
	request, err := http.NewRequest("POST", "http://authentication/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Print("error creating request to auth server", err)
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()
	// wait for the response

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("something went wrong"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Err {
		app.errorJSON(w, errors.New(jsonFromService.Msg))
		return
	}

	var payload jsonResponse
	payload.Err = false
	payload.Msg = jsonFromService.Msg
	payload.Data = jsonFromService.Data

	_ = app.writeJSON(w, http.StatusAccepted, payload)

}
