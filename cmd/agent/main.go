package main

import (
	"metrics/config/agent"
	"metrics/internal/agent/service"
	"metrics/internal/agent/storage"
)

func main() {
	agent.ParseFlagsAgent()
	st := storage.NewStorage()
	sr := service.NewService(&st)
	sr.Start()
}
