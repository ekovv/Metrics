package Service

import (
	"errors"
	"metric/internal/server/Storage"
)

type Service struct {
	storage Storage.Storage
}

func NewService(storage Storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) SetMetric(metric string, name string, value float64) error {
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
		s.storage.Inc(name, value)
	} else {
		s.storage.Set(name, value)
	}
	return nil

}
