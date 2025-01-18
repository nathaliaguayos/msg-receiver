package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const EnvPrefix = "MSG_RECEIVER"

type Config struct {
	ServiceName string `split_words:"true" default:"msg-receiver"`
	LogLevel    string `split_words:"true" default:"info"`
}

func Get() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process(EnvPrefix, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func New() *Config {
	err := godotenv.Load() // do not forget to add an .env file
	cfg, err := Get()
	if err != nil {
		panic(fmt.Errorf("invalid value(s) retrieved from environment %w", err))
	}
	return cfg
}
