package api

import (
	"github.com/grafana/grafana/pkg/registry"
)

type APIService struct {
}

func init() {
	registry.Register(&registry.Descriptor{
		Name:     "ApiService",
		Instance: &APIService{},
	})
}

func (consumer *APIService) Init() error {
	return nil
}

func (consumer *APIService) GetName() string {
	return "test"
}
