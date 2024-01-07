package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"errors"
	"keeper/utils"
	"os"
)

type Properties struct {
	DB   DBProperties `json:"db"`
	Port string       `json:"port"`
	Mode string       `json:"mode"`
}

type DBProperties struct {
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	DBName     string `json:"dbname"`
	DBHost     string `json:"dbhost"`
	DBPort     string `json:"dbport"`
}

var serverProperties *Properties = nil

// Возвращает значения содержащиеся в файле конфигурации сервера.
func GetServerProperties() Properties {
	if serverProperties == nil {
		return readAndSaveProperties()
	} else {
		return *serverProperties
	}
}

func GetServerSecret() (*ecdsa.PrivateKey, error) {
	secret, exist := os.LookupEnv("serverSecret")
	if !exist {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, errors.New("Can't generate server secret")
		}
		keyBytes, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			return nil, errors.New("Can't save server secret")
		}
		os.Setenv("serverSecret", string(keyBytes))
		return privateKey, nil
	}
	privateKey, err := x509.ParseECPrivateKey([]byte(secret))
	return privateKey, err
}

func readAndSaveProperties() Properties {
	data, err := os.ReadFile("server.properties.json")
	utils.HandleErrorWithPanic(err)
	unmarshalerr := json.Unmarshal(data, &serverProperties)
	utils.HandleErrorWithPanic(unmarshalerr)
	return *serverProperties
}