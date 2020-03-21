package config

// Config contains all of the configuration variables needed for running Melon.
// They are saved in the config.json file.
type Config struct {
	Port int `json:"port"`
}

// NewConfig creates a new Config with default values.
func NewConfig() Config {
	cfg := Config{}
	cfg.Port = 25565
	return cfg
}
