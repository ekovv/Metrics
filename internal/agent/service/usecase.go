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

func NewService(array storage.Storage) *Service {
	return &Service{array: array}
}

func (a *Service) send() error {
	//myMap := a.array.Inc()
	//for key, value := range myMap {
	//	if key == "Pollcount" {
	//		req, err := http.NewRequest("POST", "http://localhost:8080/update/counter/"+key+"/"+value, nil)
	//	}
	//	req, err := http.NewRequest("POST", "http://localhost:8080/update/gauge/"+key+"/"+value, nil)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	return nil
}
