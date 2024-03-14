package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

// Config struct holds the configuration parsed from a 'config.toml' file
type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	}
}

func main() {
	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Panicf(
			"Failed decode configuration from configuration file `config.toml`: %s",
			err.Error(),
		)
	}

	log.Printf(
		"Server started with configuration: host=\"%s\", port=\"%d\"",
		config.Server.Host,
		config.Server.Port,
	)

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and the handlers
	mux.Handle("GET /", &DefaultHandler{})

	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), mux)
}
