package main

import (
	"log"
	"os"

	"github.com/fcole90/quack/model"
	"github.com/fcole90/quack/services"
)

func main() {
	conf, err := model.ReadConfig()
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
