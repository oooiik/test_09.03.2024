package main

import (
	"context"
	"github.com/oooiik/test_09.03.2024/internal/app"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		<-signalChan
		cancel()
		time.Sleep(time.Second)
		<-ctx.Done()
	}()

	go app.NewApp().ServerRun(ctx)
}
