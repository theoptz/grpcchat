package sigctx

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ctx := New(context.Background())

	syscall.Kill(os.Getpid(), syscall.SIGINT)

	select {
	case <-ctx.Done():

	case <-time.After(500 * time.Millisecond):
		t.Errorf("Deadline must be exceeded")
	}
}
