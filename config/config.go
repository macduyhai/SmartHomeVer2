package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       string `env:"PORT" default:"9090"`
	MySQLURL   string `env:"MYSQL_URL,required"`
	APIKey     string `env:"API_KEY,required"`
	SecretKey  string `env:"SECRET_KEY,required"`
	PublicKey  string `env:"PUBLICKEY,required"`
	PrivateKey string `env:"PRIVATEKEY,required"`
}

var Conf Config

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Config {
	err := godotenv.Load(files...) // Loading config from env file
	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}
	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)

	}
	Conf = cfg
	return &cfg
}
