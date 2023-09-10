package domains

//go:generate go run github.com/vektra/mockery/v3 --name=Repository
type Repository interface {
	Set(name string, value float64)
	Inc(name string, value float64)
	Get() map[string]float64
	GetOne(name string) (float64, error)
}
