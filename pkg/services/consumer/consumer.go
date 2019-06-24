package consumer

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/registry"
)

type ConsumerService struct {
	APIService models.API `inject:""`
}

func init() {
	registry.RegisterService(&ConsumerService{})
}

func (consumer *ConsumerService) Init() error {
	name := consumer.APIService.GetName()
	spew.Dump(name)

	return nil
}
