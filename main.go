package main

import (
	"lla/api"
	"lla/config"
	"log"
)

func main() {
	config, error := config.LoadConfig()
	if error != nil {
		log.Fatal("cannot load config: ", error)
	}

	runRestfulServer(config)
}

func runRestfulServer(config *config.Config) {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
