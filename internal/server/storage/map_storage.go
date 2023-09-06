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
	//ms := make(map[string]float64)

	//myStor := &storage{
	//	m: &ms,
	//}
	//myStor.m[name] = value

	s.m[name] = value
	fmt.Println(s.m)

}

func (s *Storage) Inc(name string, value float64) {
	_, ok := s.m[name]
	new_value := 0.0
	if ok {
		new_value = s.m[name] + value
		s.m[name] = new_value
	} else {
		s.m[name] = value
	}
	fmt.Println(s.m)
}
