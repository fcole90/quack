package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const PKG_NAME = "model"

const CONFIG_PATH = "./config.json"

type Configuration struct {
	Token  string `json:"token"`
	Domain string `json:"domain"`
}

func ReadConfig(config_path ...string) (Configuration, error) {
	var config_path_inner string
	var conf Configuration

	if len(config_path) == 0 {
		config_path_inner = CONFIG_PATH
	} else if len(config_path) == 1 {
		config_path_inner = config_path[0]
	} else {
		return conf, fmt.Errorf("%s: at most 1 argument is accepted, received %d", PKG_NAME, len(config_path))
	}
	content, err := ioutil.ReadFile(config_path_inner)
	json.Unmarshal(content, &conf)
	return conf, err
}