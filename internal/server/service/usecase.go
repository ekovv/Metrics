package service

import (
	"errors"
	"metrics/internal/server/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(s storage.Storage) Service {
	return Service{
		storage: storage.NewStorage()}
}

func (s *Service) SetMetric(metric string, name string, value float64) error {

	if metric != "gauge" && metric != "counter" {
		err := errors.New("Invalid metric")
		if err != nil {
			return err
		}
	}
	if metric == "counter" {
		if value != float64(int(value)) {
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
