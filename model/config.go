package model

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type config struct {
	DSN  string
	Addr string
}

// Config 服务器配置
var Config config

func init() {
	_, err := toml.DecodeFile("config.toml", &Config)
	if err != nil {
		log.Fatal("cannot load config")
	}
}
