package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gangjun06/iduslab/models"
)

var configLocal *models.Config

// LoadConfig And Save to global Variable
func LoadConfig() (models.Config, error) {
	jsonFile, errFailedToReadConfig := ioutil.ReadFile("config.json")
	if errFailedToReadConfig != nil {
		fmt.Println(errFailedToReadConfig)
	}
	fmt.Println("Successfully Opened config.json")
	var config models.Config
	errFailedToReadConfig = json.Unmarshal(jsonFile, &config)
	configLocal = &config
	return config, errFailedToReadConfig
}

// Config Get Anywhere
func Config() *models.Config {
	return configLocal
}
