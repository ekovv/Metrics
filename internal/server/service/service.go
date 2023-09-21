package service

import (
	"fmt"
	"math"
	"metrics/internal/server/constants"
	"metrics/internal/server/domains"
)

type Service struct {
	storage domains.Repository
}

func NewService(s domains.Repository) Service {
	return Service{
		storage: s}
}

var (
	ErrInvalidMetric   = fmt.Errorf("invalid metric")
	ErrInvalidValue    = fmt.Errorf("invalid value")
	ErrInvalidDataType = fmt.Errorf("invalid data type")
)

func (s *Service) SetMetric(metric string, name string, value float64) error {
	if metric != constants.Gauge && metric != constants.Counter {
		return ErrInvalidMetric
	}

	if math.IsNaN(value) {
		return ErrInvalidValue
	}

	if metric == constants.Counter {
		if value != float64(int(value)) {
			return ErrInvalidDataType
		}
		s.storage.Inc(name, value)
	} else {
		s.storage.Set(name, value)
	}
	return nil

}

func (s *Service) GetAllMetrics() map[string]float64 {
	return s.storage.Get()
}

func (s *Service) GetVal(name string) (float64, error) {
	val, err := s.storage.GetOne(name)
	if err != nil {
		fmt.Println(name)
		return 0, err
	}
	return val, nil
}
