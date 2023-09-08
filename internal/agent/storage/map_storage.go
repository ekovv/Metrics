package storage

type Storage struct {
	metrics map[string]float64
}

type Inter interface {
	Set(metric string)
	Inc(metric string)
}

func NewStorage() Storage {
	return Storage{metrics: make(map[string]float64)}
}

func (s *Storage) Set(metric string, value float64) {
	s.metrics[metric] = value
}

func (s *Storage) Inc() map[string]float64 {
	return s.metrics
}
