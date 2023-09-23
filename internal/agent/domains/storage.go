package domains

//go:generate mockgen -source=storage.go -destination=mocks/storage.go -package=mocks
type Storage interface {
	SetGauge(metric string, value float64)
	IncCounter(metric string, value int64)
	GetGauge() map[string]float64
	GetCounter() map[string]int64
	Clear()
}
