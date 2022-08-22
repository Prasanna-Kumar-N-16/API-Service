package config

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"strings"
)

type (
	ConfigStruct struct {
		Test string `json:"test"`
		// Auth        AuthStruct  `json:"AuthStruct"`
		// DBConfig    DBConfig    `json:"dbConfig"`
		// HTTPService HTTPService `json:"httpService"`
	}
)

func Parse() (*ConfigStruct, error) {
	configFlag := flag.String("config", "config.json", "Used to read the config file")
	flag.Parse()
	return getConfig(*configFlag)
}

func getConfig(configFlag string) (*ConfigStruct, error) {
	config := ConfigStruct{}
	if strings.Contains(configFlag, ".json") {
		fileBytes, err := ioutil.ReadFile(configFlag)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(fileBytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	}
	return nil, errors.New("config file is not json file")
}
