package main

import (
	"net/http"
)




func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Err: false,
		Msg: "from broker",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}