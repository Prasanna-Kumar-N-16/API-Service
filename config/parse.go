package config

import (
	"api-service/models"
	"api-service/utils"
	"encoding/json"
	"errors"
	"flag"
	"os"
	"strings"
)

type (
	ConfigStruct struct {
		HttpConfig    *Service          `json:"httpConfig"`
		Auth          AuthStruct        `json:"authStruct"`
		MongoDBConfig models.DBConfig   `json:"mongoDBConfig"`
		PostgresQL    models.PostgresQL `json:"postgresQL"`
		Domain        string            `json:"domain"`
		EncryptKey    string            `json:"encryptKey"`
		Email         utils.EmailConfig `json:"email"`
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
	configFlag := flag.String("c", "config-dev.json", "Used to read the config file")
	flag.Parse()
	return getConfig(*configFlag)
}

func getConfig(configFlag string) (*ConfigStruct, error) {
	config := ConfigStruct{}
	if strings.Contains(configFlag, ".json") {
		fileBytes, err := os.ReadFile(configFlag)
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
	if config.EncryptKey == "" {
		return nil, errors.New("EncryptKey value is empty")
	}
	if config.Email.Username == "" || config.Email.Password == "" || strings.HasSuffix(config.Email.Username, "@gmail.com") {
		return nil, errors.New("Email Username / Password value is empty")
	}
	return &config, nil
}
