package server

import (
	"encoding/json"
	"fmt"
	"keeper/utils"
	"os"
)

type Properties struct {
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	DBName     string `json:"dbname"`
	DBHost     string `json:"dbhost"`
	DBPort     string `json:"dbport"`
	Mode       string `json:"mode"`
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

func readAndSaveProperties() Properties {
	data, err := os.ReadFile("server.properties.json")
	utils.HandleErrorWithPanic(err)
	unmarshalerr := json.Unmarshal(data, &serverProperties)
	utils.HandleErrorWithPanic(unmarshalerr)
	fmt.Println(*serverProperties)
	return *serverProperties
}
