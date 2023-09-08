package main

import (
	"metrics/internal/agent/service"
	"metrics/internal/agent/storage"
)

func main() {
	a := storage.NewStorage()
	b := service.NewService(a)
	b.Send()
	b.Update()
}
