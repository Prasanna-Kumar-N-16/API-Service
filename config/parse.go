package config

import (
	"api-service/mongodb"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"strings"
)

type (
	ConfigStruct struct {
		HttpConfig    *Service         `json:"httpConfig"`
		Auth          AuthStruct       `json:"authStruct"`
		MongoDBConfig mongodb.DBConfig `json:"mongoDBConfig"`
	}
	AuthStruct struct {
		Key string `json:"key"`
	}
	Service struct {
		Host               string     `json:"host"`
		ISSecureConnection bool       `json:"isSecureConnection" `
		SSLConfig          *SSLConfig `json:"sslConfig" `
	}
	SSLConfig struct {
		PrivateKey string `json:"privateKey"`
		CrtFile    string `json:"crtfile"`
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
		return config.validateConfig()
	}
	return nil, errors.New("config file is not json file")
}

func (config ConfigStruct) validateConfig() (*ConfigStruct, error) {
	return &config, nil
}
