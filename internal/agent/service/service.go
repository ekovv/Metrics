package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"metrics/config/agent"
	"metrics/internal/agent/domains"
	"metrics/internal/agent/models"
	"net/http"
	"runtime"
	"time"
)

type Service struct {
	storage domains.Storage
	client  domains.Client
	config  agent.Config
}

func NewService(array domains.Storage, config agent.Config) *Service {
	return &Service{
		storage: array,
		client:  http.DefaultClient,
		config:  config,
	}
}

var (
	ErrInvalidRequest = errors.New("invalid request")
)

func (a *Service) Send() error {
	myMapGauge := a.storage.GetGauge()
	myMapCounter := a.storage.GetCounter()
	var resp *http.Response

	for key, value := range myMapGauge {
		metric := models.Metrics{
			ID:    key,
			MType: "gauge",
			Delta: nil,
			Value: &value,
		}
		jsonMetric, err := json.Marshal(metric)
		if err != nil {
			fmt.Println(err)
			return err
		}
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/update/", a.config.Host), bytes.NewBuffer(jsonMetric))
		if err != nil {
			fmt.Println(err)
			return ErrInvalidRequest
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err = a.client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ErrInvalidRequest
		}
		resp.Body.Close()
		fmt.Println("отправлено гауг")
	}
	for key, value := range myMapCounter {
		metric := models.Metrics{
			ID:    key,
			MType: "counter",
			Delta: &value,
			Value: nil,
		}
		jsonMetric, err := json.Marshal(metric)
		if err != nil {
			fmt.Println(err)
			return err
		}
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/update/", a.config.Host), bytes.NewBuffer(jsonMetric))
		if err != nil {
			fmt.Println(err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err = a.client.Do(req)
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
	a.storage.IncCounter("PollCount", 1)
	a.storage.SetGauge("RandomValue", a.randomGenerate())

	fmt.Println("update")

	return nil
}

func (a *Service) Start() {
	for {
		start := time.Now()
		s := int64(a.config.ReportInterval)
		for time.Now().Unix()-start.Unix() < s {
			start := time.Now()
			err := a.Update()
			if err != nil {
				log.Fatal(err)
				return
			}
			i := time.Duration(a.config.PollInterval)
			time.Sleep(i*time.Second - time.Since(start))
		}

		err := a.Send()
		if err != nil {
			log.Fatalf("cannot send %v", err)
			return
		}

		a.storage.Clear()
	}
}

func (a *Service) randomGenerate() float64 {
	rand.NewSource(time.Now().UnixNano())
	randomFloat := rand.Float64()
	return randomFloat
}
