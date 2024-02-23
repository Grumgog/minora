package server

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"minora/utils"
	"os"

	"github.com/creasty/defaults"
)

type Properties struct {
	DB   DBProperties `json:"db"`
	Port string       `json:"port" default:"5200"`
	Host string       `json:"host" default:"localhost"`
	Mode string       `json:"mode"`
}

type DBProperties struct {
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	DBName     string `json:"dbname"`
	DBHost     string `json:"dbhost" default:"localhost"`
	DBPort     string `json:"dbport" default:"5432"`
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

func GetServerSecret() string {
	secret, exist := os.LookupEnv("serverSecret")
	if !exist {
		secret = GenerateSecretKey(64)
		err := os.Setenv("serverSecret", string(secret))
		utils.HandleErrorWithPanic(err)
		return secret
	}

	return secret
}

func GenerateSecretKey(length uint32) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

func readAndSaveProperties() Properties {
	data, err := os.ReadFile("server.properties.json")
	utils.HandleErrorWithPanic(err)
	serverProperties = &Properties{}
	defaults.Set(serverProperties)
	unmarshalerr := json.Unmarshal(data, &serverProperties)
	utils.HandleErrorWithPanic(unmarshalerr)
	return *serverProperties
}
