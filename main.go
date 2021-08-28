package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/fcole90/quack/model"
	"github.com/fcole90/quack/services"
)

const ENV_QUACK_CONFIG = "QUACK_CONFIG_DIR"
const DEFAULT_CONFIG_FILENAME = "./config.json"

func getConfigPath() string {
	envConfigFilePath := os.Getenv(ENV_QUACK_CONFIG)
	if envConfigFilePath != "" {
		return path.Join(envConfigFilePath, model.CONFIG_FILENAME)
	} else {
		return path.Join(".", model.CONFIG_FILENAME)
	}
}

func parseStringArg(fullArg string, flagArg string) string {
	return fullArg[len(flagArg):]
}

func parseIntArg(fullArg string, flagArg string) int {
	val, err := strconv.Atoi(parseStringArg(fullArg, flagArg))
	if err != nil {
		panic(err)
	}
	return val
}

func run() {
	configFilePath := getConfigPath()
	conf, err := model.ReadConfig(configFilePath)
	if err != nil {
		log.Println(err)
		log.Fatalf(`Configuration error: the configuration at "%s" is missing or unreadable`, configFilePath)
	}

	for {
		log.Println("Updating IP address...")
		res, err := services.Update(conf.Domain, conf.Token)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Response:", res)
		time.Sleep(time.Duration(conf.TimeInterval) * time.Second)
	}
}

func setConfig(args []string) {
	configFilePath := getConfigPath()
	conf, err := model.ReadConfig(configFilePath)
	if err != nil {
		fmt.Println("No previous configuration found")
	}

	tokenArg := "-token="
	domainArg := "-domain="
	timeIntervalArg := "-timeinterval="

	for _, arg := range args {
		if strings.HasPrefix(arg, tokenArg) {
			conf.Token = parseStringArg(arg, tokenArg)
		} else if strings.HasPrefix(arg, domainArg) {
			conf.Domain = parseStringArg(arg, domainArg)
		} else if strings.HasPrefix(arg, timeIntervalArg) {
			if parseStringArg(arg, timeIntervalArg) == "" {
				conf.TimeInterval = model.DEFAULT_TIME_INTERVAL
			} else {
				conf.TimeInterval = parseIntArg(arg, timeIntervalArg)
			}
		} else {
			panic(fmt.Errorf(`argument "%s" not understood`, arg))
		}
	}

	err = model.WriteConfig(conf, configFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Configuration updated")
}

func config() {
	var token string
	var domain string
	var timeInterval int
	configFilePath := getConfigPath()

	fmt.Println("Configuring Quack")
	fmt.Print("token: ")
	fmt.Scanf("%s", &token)
	fmt.Print("domain: ")
	fmt.Scanf("%s", &domain)
	fmt.Printf("update interval (seconds) (default %ds): ", model.DEFAULT_TIME_INTERVAL)
	fmt.Scanf("%d", &timeInterval)

	if timeInterval < 0 {
		fmt.Println("Invalid update interval")
		os.Exit(1)
	} else if timeInterval == 0 {
		timeInterval = model.DEFAULT_TIME_INTERVAL
	}

	err := model.WriteConfig(
		model.Configuration{
			Domain:       domain,
			Token:        token,
			TimeInterval: timeInterval,
		},
		configFilePath,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Configuration completed with success!")
}

func main() {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "config":
		config()
	case "set":
		setConfig(os.Args[2:])
	case "":
		run()
	default:
		fmt.Printf("Arguments like \"%s\" are not understood by this program\n", strings.Join(os.Args[1:], " "))
	}
}
