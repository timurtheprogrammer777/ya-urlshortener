package config

import "flag"

type ConfigType struct {
	a string
	b string
}

var serverConfig = ConfigType{
	a: "localhost:8888",
	b: "http://localhost:8000/qsd54gFg",
}

func ParseFlags() {
	flag.StringVar(&serverConfig.a, "a", "localhost:8888", "address to run server")
	flag.StringVar(&serverConfig.b, "b", "http://localhost:8000/qsd54gFg", "address to run server")

	flag.Parse()
}
