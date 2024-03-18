package main

import (
	"context"
	"github.com/oooiik/test_09.03.2024/internal/app"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		<-signalChan
		cancel()
		<-ctx.Done()
	}()

	go app.NewApp().ServerRun(ctx)
}
