package sigctx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// New returns context.Context which would cancel
// when SIGINT or SIGTERM signal receive
func New(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	return ctx
}
