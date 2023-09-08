package storage

type Storage struct {
	metrics []string
}

type Inter interface {
	Set(metric string)
}

func NewStorage(metrics []string) *Storage {
	return &Storage{metrics: make([]string, 0)}
}

func (s *Storage) Set(metric string) {

}
