package api

import "github.com/virtsevda/StandartWebServer/storage"

//General instance for API server of REST application
type Config struct{
	//Port
	BindAddr string `toml:"bind_addr"`

	//Logger level
	LoggerLevel string `toml:"logger_level"`

	//Storage
	Storage *storage.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":8080",
		LoggerLevel: "debug",
		Storage: storage.NewConfig(),
	}
}