package service

import (
	"metrics/internal/agent/storage"
)

//var memStats runtime.MemStats
//runtime.ReadMemStats(&memStats)
//
//fmt.Printf("Общий использованный объем памяти: %d bytes\n", memStats.Alloc)

type Service struct {
	array storage.Storage
}

func (a *Service) send() error {
	//
	//time.Sleep(10 * time.Second)
	//req, err := http.NewRequest("POST", "http://example.com/api/"+a.array, nil)
	return nil
}
