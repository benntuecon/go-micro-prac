package main

import (
	"encoding/json"
	"net/http"
)


type jsonResponse struct {
	Err bool `json:"err"`
	Msg string `json:"msg"`
	Data any `json:data,omitempty`

}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Err: false,
		Msg: "from broker",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}