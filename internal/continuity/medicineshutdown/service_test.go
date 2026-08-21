package medicineshutdown_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"carecontinuity/internal/continuity/medicineshutdown"
)

func TestMedicineWorkerShutdownPublicBehavior(t *testing.T) {
	coordinator := medicineshutdown.NewCoordinator()
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- coordinator.Run(ctx, func(callCtx context.Context) error { close(started); <-callCtx.Done(); return callCtx.Err() })
	}()
	<-started
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected canceled result, got %v", err)
		}
	case <-time.After(250 * time.Millisecond):
		t.Fatal("canceled operation did not stop")
	}
	if err := coordinator.Run(context.Background(), func(context.Context) error { return nil }); err != nil {
		t.Fatalf("normal operation failed: %v", err)
	}
}
