package service

import (
	"errors"
	"reflect"
)

type Service struct {
	storage storage
}

type storage interface {
	Set(name string, value float64)
	Inc(name string, value float64)
}

func NewService(s storage) Service {
	return Service{
		storage: s}
}

func (s *Service) SetMetric(metric string, name string, value float64) error {
	if metric != "gauge" && metric != "counter" {
		err := errors.New("invalid metric")
		if err != nil {
			return err
		}
	}

	if reflect.TypeOf(value).Kind() != reflect.Float64 && reflect.TypeOf(value).Kind() != reflect.Int || value == 0 {
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
