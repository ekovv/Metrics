package server

import "flag"

var FlagRunAddr string

func ParseFlagsServer() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.Parse()
}