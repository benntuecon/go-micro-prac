package main

import (
	"context"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = ":80"
	rpcPort  = ":5001"
	gRpcPort = ":50001"
	mongoUrl = "mongodb://mongo:27017"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {

	// connect to mongo client to have an access the client interface
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatal(err)
	}

	client = mongoClient

	// create a context with time out to handle the disconnection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	go app.serve()

}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Print("log server fail to start ", err)
		log.Fatal(err)
	}

}

func connectToMongo() (*mongo.Client, error) {
	// build the connection options
	clientOption, err := options.Client().ApplyURI(mongoUrl)
	clientOption.SetAuth(options.Credential{
		Username: "admin",
		Password: "Password",
	})

	// connect
	ctx := context.Background()
	c, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Print("error connecting to mongo ", err)
		return nil, err
	}

	return c, nil
}
