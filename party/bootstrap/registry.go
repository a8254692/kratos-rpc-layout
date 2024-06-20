package bootstrap

import (
	"errors"
	"github.com/go-kratos/kratos/v2/registry"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/env"
)

// NewRegistry ...
func NewRegistry() (reg registry.Registrar, err error) {
	switch env.GetMode() {
	case env.ModeDevelop, env.ModeTest, env.ModeProduction:
		//reg, err = registryx.NewKubeRegistry()
	case env.ModeLocal:
	default:
		return nil, errors.New("unknown mode")
	}
	return
}
