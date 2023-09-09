package service

import (
	"fmt"
	"math"
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
	if metric != "gauge" && metric != "counter" {
		return ErrInvalidMetric
	}

	if math.IsNaN(value) || value == 0 {
		return ErrInvalidValue
	}

	if metric == "counter" {
		if value != float64(int(value)) {
			return ErrInvalidDataType
		}
		s.storage.Inc(name, value)
	} else {
		s.storage.Set(name, value)
	}
	return nil

}
