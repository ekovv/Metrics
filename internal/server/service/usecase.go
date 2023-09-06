package service

import (
	"errors"
	"metrics/internal/server/my_storage"
)

type Service struct {
	storage my_storage.Storage
}

func NewService(s my_storage.Storage) Service {
	return Service{
		storage: my_storage.NewStorage()}
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
