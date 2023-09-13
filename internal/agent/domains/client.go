package domains

import "net/http"

//go:generate mockgen -source=client.go -destination=mocks/client.go -package=mocks
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}
