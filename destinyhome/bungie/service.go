package bungie

import (
	"net/http"
)

type Service interface {
	Destiny2() *destiny2Service
}

type service struct {
	client   *http.Client
	apiKey   string
	basePath string

	destiny2 *destiny2Service
}

func NewService(c *http.Client, k string) (Service, error) {
	s := &service{
		client:   c,
		apiKey:   k,
		basePath: "https://www.bungie.net/Platform",
	}
	s.destiny2 = &destiny2Service{s}

	return s, nil
}

func (s *service) Destiny2() *destiny2Service {
	return s.destiny2
}
