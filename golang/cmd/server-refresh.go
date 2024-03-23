package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChanInt := make(chan os.Signal, 1)
	defer close(signalChanInt)
	signal.Notify(signalChanInt, os.Interrupt)

	signalChanKill := make(chan os.Signal, 1)
	defer close(signalChanKill)
	signal.Notify(signalChanKill, syscall.SIGQUIT)

	defer func() {
		<-signalChanKill
		fmt.Println()
	}()

	go func() {
		for {
			cmdCtx, cmdCancel := context.WithCancel(context.Background())
			defer cmdCancel()

			clearCmd := exec.Command("clear")
			clearCmd.Stdout = os.Stdout
			clearCmd.Run()

			cmd := exec.CommandContext(cmdCtx, "go", "run", "cmd/server-run.go")
			cmd.Dir = "/app"
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout

			if err := cmd.Start(); err != nil {
				fmt.Println("Error starting command:", err)
				return
			}

			if err := cmd.Wait(); err != nil {
				fmt.Println("Error waiting for command to finish:", err)
			}

			<-signalChanInt
			fmt.Println()
			time.Sleep(time.Millisecond * 200)
		}
		//signalChanKill <- nil
	}()
}
