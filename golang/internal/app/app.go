package app

import (
	"context"
	"github.com/oooiik/test_09.03.2024/internal/provider"
)

type Interface interface {
	ServerRun(ctx context.Context)
}

type app struct {
	http provider.Http
}

func NewApp() Interface {
	a := app{}
	a.initProviders()
	return &a
}

func (a *app) initProviders() {
	a.http = provider.NewHttp()
}

func (a *app) ServerRun(c context.Context) {
	a.http.ServerRun(c)
}
