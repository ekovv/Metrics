package server

import (
	"flag"
	"os"
)

var FlagRunAddr string

func ParseFlagsServer() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.Parse()
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}

}
