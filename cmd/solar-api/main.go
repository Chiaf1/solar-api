package main

import (
	"log"

	"github.com/chiaf1/solar-api/internal/config"
)

const CONFIG_PATH = "./config.yaml"

func main() {
	// Load configs from file
	var conf config.Config
	err := conf.Load(CONFIG_PATH)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Validate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config Loaded")
}
