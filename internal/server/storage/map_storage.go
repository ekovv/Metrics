package storage

import "fmt"

type Storage struct {
	m map[string]float64
}

func NewStorage() Storage {
	return Storage{m: make(map[string]float64)}
}

type Interface interface {
	set(name string, value float64)
	inc(name string, value float64)
}

func (s *Storage) Set(name string, value float64) {
	s.m[name] = value
	fmt.Println(s.m)

}

func (s *Storage) Inc(name string, value float64) {
	_, ok := s.m[name]
	newValue := 0.0
	if ok {
		newValue = s.m[name] + value
		s.m[name] = newValue
	} else {
		s.m[name] = value
	}
	fmt.Println(s.m)
}
