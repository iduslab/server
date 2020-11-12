package utils

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/iduslab/backend/models"
)

var configLocal *models.Config

// LoadConfig And Save to global Variable
func LoadConfig() {
	rawConfig, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalln("Failed to load config.")
	}

	var config models.Config
	if _, err := toml.Decode(string(rawConfig), &config); err != nil {
		log.Fatalln("Failed to parsing config.")
	}
	configLocal = &config
}

// Config Get Anywhere
func Config() *models.Config {
	return configLocal
}
