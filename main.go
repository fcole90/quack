package main

import (
	"log"
	"os"

	"github.com/fcole90/quack/model"
	"github.com/fcole90/quack/services"
)

func main() {
	config_file := os.Getenv("QUACK_CONFIG")
	if config_file == "" {
		config_file = model.CONFIG_PATH
	}

	conf, err := model.ReadConfig(config_file)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Println("Updating IP address...")
	res, err := services.Update(conf.Domain, conf.Token)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Response:", res)
}
