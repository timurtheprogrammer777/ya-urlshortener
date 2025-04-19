package config

import "flag"

type ConfigType struct {
	Address      string
	BaseShortUrl string
}

var ServerConfig = ConfigType{
	Address:      "localhost:8888",
	BaseShortUrl: "http://localhost:8000/qsd54gFg",
}

func ParseFlags() {
	flag.StringVar(&ServerConfig.Address, "a", "localhost:8888", "address to run server")
	flag.StringVar(&ServerConfig.BaseShortUrl, "b", "http://localhost:8000/qsd54gFg", "address to run server")

	flag.Parse()
}
