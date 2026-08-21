package medicineshutdown

import (
	"context"
	"fmt"
)

type Coordinator struct{ worker *DeliveryWorker }

func NewCoordinator() *Coordinator {
	return &Coordinator{worker: NewDeliveryWorker(context.Background(), "worker")}
}

func (c *Coordinator) Run(ctx context.Context, operation func(context.Context) error) error {
	callCtx, err := c.worker.Scope(ctx)
	if err != nil {
		return err
	}
	if err := operation(callCtx); err != nil {
		return fmt.Errorf("medicine worker shutdown operation: %w", err)
	}
	return nil
}
