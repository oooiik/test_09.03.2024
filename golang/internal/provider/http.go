package provider

import (
	"context"
	"errors"
	"github.com/oooiik/test_09.03.2024/internal/config"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"net/http"
	"sync"
)

type Http interface {
	ServerRun(ctx context.Context)
}

type httpProvider struct {
}

func NewHttp() Http {
	return &httpProvider{}
}

func (h *httpProvider) ServerRun(c context.Context) {
	srv := &http.Server{
		Addr: config.Load().Server.Adders(),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		logger.Info("Starting listen server", config.Load().Server.Adders())

		err := srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Info(err)
			} else {
				logger.Fatal(err)
			}
		}
	}()

	<-c.Done()
	err := srv.Shutdown(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Shutdown listen server!")
	wg.Wait()
}
