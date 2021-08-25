package model

import (
	"encoding/json"
	"io/ioutil"
)

const PKG_NAME = "model"

type Configuration struct {
	Token  string `json:"token"`
	Domain string `json:"domain"`
}

func ReadConfig(config_path string) (Configuration, error) {
	var conf Configuration
	content, err := ioutil.ReadFile(config_path)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(content, &conf)
	return conf, err

}

func WriteConfig(conf Configuration, config_path string) error {
	content, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config_path, content, 0600)
	return err
}
