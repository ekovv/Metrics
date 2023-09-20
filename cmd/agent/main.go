package main

import (
	"metrics/config/agent"
	"metrics/internal/agent/service"
	"metrics/internal/agent/storage"
)

func main() {
	config := agent.New()
	st := storage.NewStorage()
	sr := service.NewService(&st, config)
	sr.Start()
}
