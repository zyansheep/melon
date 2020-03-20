package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config is the Melon configuration struct.
type Config struct {
	MelonAddress string `json:"melonAddress"`
	Name         string `json:"name"`
	HostAddress  string `json:"hostAddress"`
	HostPort     int    `json:"hostPort"`
}

// NewConfig creates a new Config with default values.
func NewConfig() Config {
	config := Config{}
	config.MelonAddress = "0.0.0.0:19132"
	config.Name = ColorDarkGreen + FormatBold + "Melon Proxy Server"
	config.HostAddress = "example.com"
	config.HostPort = 25565
	return config
}

// LoadConfig loads the Config from the config.json file, or create one if config.json doesn't exist.
func LoadConfig() Config {
	config := Config{}

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Creating a default config.json file.")

		config = NewConfig()

		file, err = json.MarshalIndent(config, "", "    ")
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("config.json", file, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		err = json.Unmarshal([]byte(file), &config)
		if err != nil {
			panic(err)
		}
	}

	return config
}
