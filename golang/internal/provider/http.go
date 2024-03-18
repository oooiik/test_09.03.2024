package provider

import (
	"context"
	"errors"
	"github.com/oooiik/test_09.03.2024/internal/config"
	"github.com/oooiik/test_09.03.2024/internal/http/controller"
	"github.com/oooiik/test_09.03.2024/internal/http/router"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"net/http"
	"sync"
)

type Http interface {
	ServerRun(ctx context.Context)
}

type httpProvider struct {
	router      *router.Router
	controllers []*controller.Interface
	server      *http.Server
}

func NewHttp() Http {
	h := httpProvider{}
	h.initController()
	h.initRouter()
	h.initServer()
	return &h
}

func (h *httpProvider) ServerRun(c context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		logger.Info("Starting listen server", config.Load().Server.Adders())

		err := h.server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Info(err)
			} else {
				logger.Fatal(err)
			}
		}
	}()

	<-c.Done()
	err := h.server.Shutdown(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Shutdown listen server!")
	wg.Wait()
}

func (h *httpProvider) initController() {
	// TODO
}
func (h *httpProvider) initRouter() {
	h.router = router.New()
}
func (h *httpProvider) initServer() {
	h.server = &http.Server{
		Addr:    config.Load().Server.Adders(),
		Handler: h.router.Handler(),
	}
}
