package app

import (
	"github.com/oooiik/test_09.03.2024/internal/provider"
)

type Interface interface {
}

type app struct {
	Providers []*provider.Interface
}

var singleton Interface

func App() Interface {
	if singleton == nil {
		a := app{}
		singleton = &a
	}
	return singleton
}

func (a *app) initProviders() {
	for _, p := range provider.List {
		p.Bind(a)
		p.Boot()
		p.Register()
	}
}
