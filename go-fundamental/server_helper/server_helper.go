package serverhelper

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

// SignalContext ...
func SignalContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		glog.Infof("listening for shutdown signal")
		<-sigs
		glog.Infof("shutdown signal received")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	return ctx
}
