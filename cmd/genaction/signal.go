package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func listenSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go func() {
		<-sigChan
		cancel()
	}()

	return ctx
}
