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
	var resp *http.Response
	for key, value := range myMap {
		req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f", key, value), nil)
		resp, _ = client.Do(req)

	}
	req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/counter/%s/%f", key, value), nil)
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	fmt.Println("done")
	time.Sleep(10 * time.Second)

	return nil
}

func (a *Service) Update() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	a.array.SetGauge("Alloc", float64(memStats.Alloc))
	a.array.SetGauge("BuckHashSys", float64(memStats.BuckHashSys))
	a.array.SetGauge("Frees", float64(memStats.Frees))
	a.array.SetGauge("GCCPUFraction", float64(memStats.GCSys))
	a.array.SetGauge("HeapAlloc", float64(memStats.HeapAlloc))
	a.array.SetGauge("HeapIdle", float64(memStats.HeapIdle))
	a.array.SetGauge("HeapInuse", float64(memStats.HeapInuse))
	a.array.SetGauge("HeapObjects", float64(memStats.HeapObjects))
	a.array.SetGauge("HeapReleased", float64(memStats.HeapReleased))
	a.array.SetGauge("HeapSys", float64(memStats.HeapSys))
	a.array.SetGauge("LastGC", float64(memStats.LastGC))
	a.array.SetGauge("Lookups", float64(memStats.Lookups))
	a.array.SetGauge("MCacheInuse", float64(memStats.MCacheInuse))
	a.array.SetGauge("MCacheSys", float64(memStats.MCacheSys))
	a.array.SetGauge("MSpanInuse", float64(memStats.MSpanInuse))
	a.array.SetGauge("MSpanSys", float64(memStats.MSpanSys))
	a.array.SetGauge("Mallocs", float64(memStats.Mallocs))
	a.array.SetGauge("NextGC", float64(memStats.NextGC))
	a.array.SetGauge("NumForcedGC", float64(memStats.NumForcedGC))
	a.array.SetGauge("NumGC", float64(memStats.NumGC))
	a.array.SetGauge("OtherSys", float64(memStats.OtherSys))
	a.array.SetGauge("PauseTotalNs", float64(memStats.PauseTotalNs))
	a.array.SetGauge("StackInuse", float64(memStats.StackInuse))
	a.array.SetGauge("StackSys", float64(memStats.StackSys))
	a.array.SetGauge("Sys", float64(memStats.Sys))
	a.array.SetGauge("TotalAlloc", float64(memStats.TotalAlloc))
	a.array.SetCounter("PollCount", 0)
	a.array.SetGauge("RandomValue", 13.2)
	fmt.Println("update")
	time.Sleep(10 * time.Second)
	return nil
}

func increm
