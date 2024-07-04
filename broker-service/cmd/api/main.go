package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "localhost:3000"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker server on %s\n", webPort)

	srv := &http.Server{

		Addr:    fmt.Sprintf("%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
