package main

import (
	"metrics/internal/agent/service"
	"metrics/internal/agent/storage"
)

func main() {
	st := storage.NewStorage()
	sr := service.NewService(&st)

	//start
}
