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
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			continue

		}
		req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f", key, value), nil)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

	}

	fmt.Println("done")
	time.Sleep(10 * time.Second)

	return nil
}

func (a *Service) Update() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	a.array.Set("Alloc", float64(memStats.Alloc))
	a.array.Set("BuckHashSys", float64(memStats.BuckHashSys))
	a.array.Set("Frees", float64(memStats.Frees))
	a.array.Set("GCCPUFraction", float64(memStats.GCSys))
	a.array.Set("HeapAlloc", float64(memStats.HeapAlloc))
	a.array.Set("HeapIdle", float64(memStats.HeapIdle))
	a.array.Set("HeapInuse", float64(memStats.HeapInuse))
	a.array.Set("HeapObjects", float64(memStats.HeapObjects))
	a.array.Set("HeapReleased", float64(memStats.HeapReleased))
	a.array.Set("HeapSys", float64(memStats.HeapSys))
	a.array.Set("LastGC", float64(memStats.LastGC))
	a.array.Set("Lookups", float64(memStats.Lookups))
	a.array.Set("MCacheInuse", float64(memStats.MCacheInuse))
	a.array.Set("MCacheSys", float64(memStats.MCacheSys))
	a.array.Set("MSpanInuse", float64(memStats.MSpanInuse))
	a.array.Set("MSpanSys", float64(memStats.MSpanSys))
	a.array.Set("Mallocs", float64(memStats.Mallocs))
	a.array.Set("NextGC", float64(memStats.NextGC))
	a.array.Set("NumForcedGC", float64(memStats.NumForcedGC))
	a.array.Set("NumGC", float64(memStats.NumGC))
	a.array.Set("OtherSys", float64(memStats.OtherSys))
	a.array.Set("PauseTotalNs", float64(memStats.PauseTotalNs))
	a.array.Set("StackInuse", float64(memStats.StackInuse))
	a.array.Set("StackSys", float64(memStats.StackSys))
	a.array.Set("Sys", float64(memStats.Sys))
	a.array.Set("TotalAlloc", float64(memStats.TotalAlloc))
	a.array.Set("PollCount", 0)
	a.array.Set("RandomValue", 13.2)
	fmt.Println("update")
	time.Sleep(10 * time.Second)
	return nil
}
