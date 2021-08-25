package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fcole90/quack/model"
	"github.com/fcole90/quack/services"
)

const ENV_QUACK_CONFIG = "QUACK_CONFIG"
const DEFAULT_CONFIG_FILENAME = "config.json"

func run() {
	tokenArgPtr := flag.String("token", "", "token for Duck DNS")
	domainArgPtr := flag.String("domain", "", "domain name to update on Duck DNS")
	flag.Parse()
	config_file := os.Getenv(ENV_QUACK_CONFIG)

	var conf model.Configuration
	if *tokenArgPtr != "" && *domainArgPtr != "" {
		conf = model.Configuration{
			Domain: *domainArgPtr,
			Token:  *tokenArgPtr,
		}
	} else if config_file != "" {
		var err error
		conf, err = model.ReadConfig(config_file)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("you need to either specify a config file or provide the required arguments"))
	}

	log.Println("Updating IP address...")
	res, err := services.Update(conf.Domain, conf.Token)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Response:", res)
}

func config() {
	var token string
	var domain string
	var config_dir string
	config_file_default := os.Getenv(ENV_QUACK_CONFIG)

	fmt.Println("Configuring Quack:")
	fmt.Print("token: ")
	fmt.Scanf("%s", &token)
	fmt.Print("domain: ")
	fmt.Scanf("%s", &domain)

	var message = "config directory: "
	if config_file_default != "" {
		message = fmt.Sprintf(`%s(default: "%s") `, message, config_file_default)
	}
	fmt.Print(message)
	fmt.Scanf("%s", &config_dir)

	err := model.WriteConfig(
		model.Configuration{
			Domain: domain,
			Token:  token,
		},
		filepath.Join(config_dir, DEFAULT_CONFIG_FILENAME),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Configured!")
}

func main() {
	command := os.Args[1]

	switch command {
	case "config":
		config()
	default:
		run()
	}
}
