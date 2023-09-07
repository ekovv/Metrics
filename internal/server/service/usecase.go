package service

import (
	"errors"
	"metrics/internal/server/storage"
	"reflect"
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
		err := errors.New("invalid metric")
		if err != nil {
			return err
		}
	}

	if reflect.TypeOf(value).Kind() != reflect.Float64 || reflect.TypeOf(value).Kind() != reflect.Int {
		err := errors.New("invalid value type")
		if err != nil {
			return err
		}
	}

	if metric == "counter" {
		if value != float64(int(value)) {
			err := errors.New("invalid Data Type")
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
