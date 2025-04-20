package config

import "flag"

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
	flag.StringVar(&ServerConfig.BaseShortURL, "b", "http://localhost:8000/qsd54gFg", "address to run server")

	flag.Parse()
}
