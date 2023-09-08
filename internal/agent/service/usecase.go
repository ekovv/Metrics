package service

import (
	"fmt"
	"metrics/internal/agent/storage"
	"net/http"
	"runtime"
	"time"
)

type Service struct {
	array storage.Storage
}

func NewService(array storage.Storage) Service {
	return Service{array: array}
}

func (a *Service) Send() error {
	myMap := a.array.Inc()
	client := &http.Client{}
	for key, value := range myMap {
		if key == "Pollcount" {
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/counter/%s/%f", key, value), nil)
			_, err := client.Do(req)
			if err != nil {
				return err
			}
			continue
		}
		req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f", key, value), nil)
		_, err := client.Do(req)
		if err != nil {
			return err
		}

	}
	fmt.Println("done")
	time.Sleep(10 * time.Second)

	return nil
}

func (s *Service) Update() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	s.array.Set("Alloc", float64(memStats.Alloc))
	s.array.Set("BuckHashSys", float64(memStats.BuckHashSys))
	s.array.Set("Frees", float64(memStats.Frees))
	s.array.Set("GCCPUFraction", float64(memStats.GCSys))
	s.array.Set("HeapAlloc", float64(memStats.HeapAlloc))
	s.array.Set("HeapIdle", float64(memStats.HeapIdle))
	s.array.Set("HeapInuse", float64(memStats.HeapInuse))
	s.array.Set("HeapObjects", float64(memStats.HeapObjects))
	s.array.Set("HeapReleased", float64(memStats.HeapReleased))
	s.array.Set("HeapSys", float64(memStats.HeapSys))
	s.array.Set("LastGC", float64(memStats.LastGC))
	s.array.Set("Lookups", float64(memStats.Lookups))
	s.array.Set("MCacheInuse", float64(memStats.MCacheInuse))
	s.array.Set("MCacheSys", float64(memStats.MCacheSys))
	s.array.Set("MSpanInuse", float64(memStats.MSpanInuse))
	s.array.Set("MSpanSys", float64(memStats.MSpanSys))
	s.array.Set("Mallocs", float64(memStats.Mallocs))
	s.array.Set("NextGC", float64(memStats.NextGC))
	s.array.Set("NumForcedGC", float64(memStats.NumForcedGC))
	s.array.Set("NumGC", float64(memStats.NumGC))
	s.array.Set("OtherSys", float64(memStats.OtherSys))
	s.array.Set("PauseTotalNs", float64(memStats.PauseTotalNs))
	s.array.Set("StackInuse", float64(memStats.StackInuse))
	s.array.Set("StackSys", float64(memStats.StackSys))
	s.array.Set("Sys", float64(memStats.Sys))
	s.array.Set("TotalAlloc", float64(memStats.TotalAlloc))
	s.array.Set("PollCount", 0)
	s.array.Set("RandomValue", 13.2)
	fmt.Println("update")
	time.Sleep(10 * time.Second)
	return nil
}
