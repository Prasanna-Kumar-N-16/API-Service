package config

import (
	"api-service/models"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"strings"
)

type (
	ConfigStruct struct {
		HttpConfig    *Service        `json:"httpConfig"`
		Auth          AuthStruct      `json:"authStruct"`
		MongoDBConfig models.DBConfig `json:"mongoDBConfig"`
		Domain        string          `json:"domain"`
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
	configFlag := flag.String("config", "config-dev.json", "Used to read the config file")
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
	if config.HttpConfig.Host == "" {
		return nil, errors.New("host value is empty")
	}
	if config.HttpConfig.ISSecureConnection &&
		(config.HttpConfig.SSLConfig.CrtFile == "" || config.HttpConfig.SSLConfig.PrivateKey == "") {
		return nil, errors.New("CrtFile/PrivateKey filepath not defined in Secure connection")
	}
	if config.Domain == "" {
		return nil, errors.New("admin domain value is empty")
	}
	return &config, nil
}
