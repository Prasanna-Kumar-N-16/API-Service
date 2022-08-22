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
		HttpConfig *Service `json:"httpConfig" yml:"httpConfig"`
		// Auth        AuthStruct  `json:"AuthStruct"`
		// DBConfig    DBConfig    `json:"dbConfig"`
		// HTTPService HTTPService `json:"httpService"`
	}
	Service struct {
		Host               string     `json:"host" yml:"host"`
		ISSecureConnection bool       `json:"isSecureConnection" yml:"isSecureConnection"`
		SSLConfig          *SSLConfig `json:"sslConfig" yml:"sslConfig"`
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
		return &config, nil
	}
	return nil, errors.New("config file is not json file")
}
