package server

import (
	"flag"
	"os"
)

//type Flags struct {
//	FlagRunAddr *string
//}
//
//var f Flags
//
//func Config() *Flags {
//	flag.StringVar(f.FlagRunAddr, "a", ":8080", "address and port to run server")
//	flag.Parse()
//	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
//		f.FlagRunAddr = &envRunAddr
//	}
//	return &Flags{
//		FlagRunAddr: f.FlagRunAddr,
//	}
//}

type Config struct {
	Host string
}

type F struct {
	host *string
}

var f F

func init() {
	f.host = flag.String("a", "localhost:8080", "-a=host")
}

func New() *Config {
	flag.Parse()
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		f.host = &envRunAddr
	}

	return &Config{
		Host: *f.host,
	}
}
