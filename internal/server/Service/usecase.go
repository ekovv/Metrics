package Service

import (
	"errors"
	"metrics/internal/server/Storage"
)

type Service struct {
	storage Storage.Storage
}

func NewService(storage Storage.Storage) *Service {
	return &Service{storage: Storage.NewStorage(map[string]float64{})}
}

func (s *Service) SetMetric(metric string, name string, value float64) error {
	as := NewService(s.storage)

	if metric != "gauge" && metric != "counter" {
		err := errors.New("Invalid metric")
		if err != nil {
			return err
		}
	}
	if metric == "counter" {
		if float64(int(value)) != value {
			err := errors.New("Invalid Data Type")
			if err != nil {
				return err
			}
		}
		as.storage.Inc(name, value)
	} else {
		as.storage.Set(name, value)
	}
	return nil

}
