package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gangjun06/iduslab/structure"
)

var configLocal *structure.Config

// LoadConfig And Save to global Variable
func LoadConfig() (structure.Config, error) {
	jsonFile, errFailedToReadConfig := ioutil.ReadFile("config.json")
	if errFailedToReadConfig != nil {
		fmt.Println(errFailedToReadConfig)
	}
	fmt.Println("Successfully Opened config.json")
	var config structure.Config
	errFailedToReadConfig = json.Unmarshal(jsonFile, &config)
	configLocal = &config
	return config, errFailedToReadConfig
}

// Config Get Anywhere
func Config() *structure.Config {
	return configLocal
}
