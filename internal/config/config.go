package config

import (
	"flag"
	"os"
)

type ConfigType struct {
	Address      string
	BaseShortURL string
}

var ServerConfig = ConfigType{
	Address:      ":8080",
	BaseShortURL: "http://localhost:8080/qsd54gFg",
}

func ParseFlags() {
	flag.StringVar(&ServerConfig.Address, "a", ":8080", "address to run server")
	flag.StringVar(&ServerConfig.BaseShortURL, "b", "http://localhost:8080", "base short URL")

	if osAddress := os.Getenv("SERVER_ADDRESS"); osAddress != "" {
		ServerConfig.Address = osAddress
	}

	if osBaseURL := os.Getenv("BASE_URL"); osBaseURL != "" {
		ServerConfig.BaseShortURL = osBaseURL
	}

	flag.Parse()
}
