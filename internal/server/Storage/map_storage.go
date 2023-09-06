package Storage

import "fmt"

type Storage struct {
	m map[string]float64
}

func NewStorage(m map[string]float64) *Storage {
	return &Storage{m: make(map[string]float64)}
}

type StorageInterface interface {
	set(name string, value float64)
	inc(name string, value float64)
}

func (s *Storage) Set(name string, value float64) {
	//ms := make(map[string]float64)

	//myStor := &Storage{
	//	m: &ms,
	//}
	//myStor.m[name] = value
	s.m[name] = value
	fmt.Println(s.m)

}

func (s *Storage) Inc(name string, value float64) {
	_, ok := s.m[name]
	if ok {
		var v float64
		v = s.m[name] + value
		s.m[name] = v
	} else {
		s.m[name] = value
	}
	fmt.Println(s.m)
}
