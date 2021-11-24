package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Encrypt struct {
	PrivateKey string
	PublicKey  string
	KeyTitle   string
}

type GRpc struct {
	Addr string
}

type Config struct {
	Encrypt *Encrypt
	GRpc    *GRpc
}

func GetConfig() (config *Config) {
	if _, err := toml.DecodeFile("./config/config.toml", &config); err != nil {
		log.Fatal(err.Error())
	}
	return
}
