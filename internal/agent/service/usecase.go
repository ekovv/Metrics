package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

type Service struct {
	storage storage
}

func NewService(array storage) Service {
	return Service{storage: array}
}

type storage interface {
	SetGauge(metric string, value float64)
	SetCounter(metric string, value int)
	GetGauge() map[string]float64
	GetCounter() map[string]int
}

func (a *Service) Send() error {
	myMapGauge := a.storage.GetGauge()
	myMapCounter := a.storage.GetCounter()
	client := &http.Client{}
	var resp *http.Response
	for key, value := range myMapGauge {
		req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f", key, value), nil)
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Body.Close()
		fmt.Println("отправлено гауг")
	}
	for key, value := range myMapCounter {
		value += 1
		req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/update/counter/%s/%d", key, value), nil)
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Body.Close()
		fmt.Println("отправлено коунтер")
	}
	fmt.Println("done")

	return nil
}

// сделать что то респ боди клоус
func (a *Service) Update() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	a.storage.SetGauge("Alloc", float64(memStats.Alloc))
	a.storage.SetGauge("BuckHashSys", float64(memStats.BuckHashSys))
	a.storage.SetGauge("Frees", float64(memStats.Frees))
	a.storage.SetGauge("GCCPUFraction", float64(memStats.GCSys))
	a.storage.SetGauge("HeapAlloc", float64(memStats.HeapAlloc))
	a.storage.SetGauge("HeapIdle", float64(memStats.HeapIdle))
	a.storage.SetGauge("HeapInuse", float64(memStats.HeapInuse))
	a.storage.SetGauge("HeapObjects", float64(memStats.HeapObjects))
	a.storage.SetGauge("HeapReleased", float64(memStats.HeapReleased))
	a.storage.SetGauge("HeapSys", float64(memStats.HeapSys))
	a.storage.SetGauge("LastGC", float64(memStats.LastGC))
	a.storage.SetGauge("Lookups", float64(memStats.Lookups))
	a.storage.SetGauge("MCacheInuse", float64(memStats.MCacheInuse))
	a.storage.SetGauge("MCacheSys", float64(memStats.MCacheSys))
	a.storage.SetGauge("MSpanInuse", float64(memStats.MSpanInuse))
	a.storage.SetGauge("MSpanSys", float64(memStats.MSpanSys))
	a.storage.SetGauge("Mallocs", float64(memStats.Mallocs))
	a.storage.SetGauge("NextGC", float64(memStats.NextGC))
	a.storage.SetGauge("NumForcedGC", float64(memStats.NumForcedGC))
	a.storage.SetGauge("NumGC", float64(memStats.NumGC))
	a.storage.SetGauge("OtherSys", float64(memStats.OtherSys))
	a.storage.SetGauge("PauseTotalNs", float64(memStats.PauseTotalNs))
	a.storage.SetGauge("StackInuse", float64(memStats.StackInuse))
	a.storage.SetGauge("StackSys", float64(memStats.StackSys))
	a.storage.SetGauge("Sys", float64(memStats.Sys))
	a.storage.SetGauge("TotalAlloc", float64(memStats.TotalAlloc))
	a.storage.SetCounter("PollCount", 0)
	a.storage.SetGauge("RandomValue", a.randomGenerate())
	fmt.Println("update")
	return nil
}

// сделать рандом для рандомвалуе, сделать увеличивающийся счетик для поллкаунт, сделать функцию старт с таймером

func (a *Service) Start() {
	for {
		err := a.Update()
		if err != nil {
			return
		}
		//err = a.Send()
		//if err != nil {
		//	return
		//}
		a.Send()
	}
}

func (a *Service) randomGenerate() float64 {
	rand.NewSource(time.Now().UnixNano())
	randomFloat := rand.Float64()
	return randomFloat
}
