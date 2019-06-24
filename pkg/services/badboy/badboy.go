package badboy

import (
	"github.com/grafana/grafana/pkg/registry"
)

type BadBoyService struct {
}

func init() {
	registry.RegisterService(&BadBoyService{})
}

func (consumer *BadBoyService) Init() error {
	return nil
}

func (consumer *BadBoyService) GetName() string {
	return "nope"
}
