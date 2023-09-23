package storage

type Storage struct {
	metricsGauge   map[string]float64
	metricsCounter map[string]int64
}

func NewStorage() Storage {
	return Storage{
		metricsGauge:   make(map[string]float64),
		metricsCounter: make(map[string]int64),
	}
}

func (s *Storage) SetGauge(metric string, value float64) {
	s.metricsGauge[metric] = value
}

func (s *Storage) IncCounter(metric string, value int64) {
	s.metricsCounter[metric] += value
}

func (s *Storage) GetGauge() map[string]float64 { //для поллкаунт увеличивает на 1 значение
	return s.metricsGauge
}

func (s *Storage) GetCounter() map[string]int64 { //для поллкаунт увеличивает на 1 значение
	return s.metricsCounter
}

func (s *Storage) Clear() {
	s.metricsCounter = make(map[string]int64)
	s.metricsGauge = make(map[string]float64)
}
