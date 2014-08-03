package config

import (
	"code.google.com/p/gcfg"
	"log"
)

type Config struct {
	Host struct {
		Port string
		Path string
	}
}

var cfg Config
var loaded bool

func Get() Config {
	if !loaded {
		Load()
	}
	return cfg
}

func Load() {
	err := gcfg.ReadFileInto(&cfg, "/etc/pad/pad.gcfg")
	if err != nil {
		log.Printf("Failed to read config: %v", err)
	} else {
		loaded = true
	}
}
