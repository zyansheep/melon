package util

import (
    "encoding/json"
    "fmt"
	"io/ioutil"
)

// Melon configuration struct.
type Config struct {
    Address string `json:"address"`
}

// Create a new Config with default values.
func NewConfig() Config {
    config := Config{}
    config.Address = "0.0.0.0:25565"
    return config
}

// Load the Config from the config.json file, or create one if config.json doesn't exist.
func LoadConfig() Config {
    config := Config{}

    file, error := ioutil.ReadFile("config.json")
    if error != nil {
        config = NewConfig()

        file, error = json.MarshalIndent(config, "", "    ")
        if error != nil {
            panic(error)
        }

        error = ioutil.WriteFile("config.json", file, 0644)
        if error != nil {
            panic(error)
        }

        fmt.Println("Created a blank 'config.json' file.")
    } else {
        error = json.Unmarshal([]byte(file), &config)
        if error != nil {
            panic(error)
        }
    }

    return config
}
