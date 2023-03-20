package main

import (
	"log"
	"net/http"
)


const webPort = ":8000"
type Config struct {}

func main() {

	app := Config{}

	log.Println("Starting server on port ", webPort)

	srv := &http.Server{
		Addr: webPort,
		Handler: app.routes(),

	}


	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("server fail to start ",err)
	}

}