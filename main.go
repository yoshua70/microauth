package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

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
		"Listening to incoming request at http://%s:%d",
		config.Server.Host,
		config.Server.Port,
	)
}
