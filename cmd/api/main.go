package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	//set application config
	var app application

	//read from command line

	//connect to database

	//start a webserver
	app.Domain = "example.com"

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
