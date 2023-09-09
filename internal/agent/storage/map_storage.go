package storage

type Storage struct {
	metricsGauge   map[string]float64
	metricsCounter map[string]int
}

func NewStorage() Storage {
	return Storage{
		metricsGauge:   make(map[string]float64),
		metricsCounter: make(map[string]int),
	}
}

func (s *Storage) SetGauge(metric string, value float64) {
	s.metricsGauge[metric] = value
}

func (s *Storage) SetCounter(metric string, value int) {
	s.metricsCounter[metric] = value
}

func (s *Storage) GetGauge() map[string]float64 { //для поллкаунт увеличивает на 1 значение
	return s.metricsGauge
}

func (s *Storage) GetCounter() map[string]int { //для поллкаунт увеличивает на 1 значение
	return s.metricsCounter
}
