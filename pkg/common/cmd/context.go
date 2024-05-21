package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// CreateContext creates a new context with a cancel function
func CreateContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c: cancel()
		case <-ctx.Done(): return
		}
	}()
	return ctx
}