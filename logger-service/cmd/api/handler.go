package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	UserId string `json:"userId"`
	Action string `json:"action`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		UserId: requestPayload.UserId,
		Action: requestPayload.Action,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Err: false,
		Msg: "log written",
	}

	_ = app.writeJSON(w, http.StatusAccepted, resp)
}
