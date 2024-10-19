package config

import (
	"log"
)

type Config struct {
	PORT    string
	VERSION string
}

func NewDefaultConfig() *Config {
	return &Config{
		PORT:    "8989",
		VERSION: "v0.1.0",
	}
}

func (c *Config) Banner() {
	printLogo()
	log.Printf("Rover Server running")
	log.Printf("Port:\t\t%s", c.PORT)
	log.Printf("Documentation:\thttps://rover.subrotokumar.dev")
	log.Printf("Source code:\thttps://github.com/subrotokumar/rover")
}
