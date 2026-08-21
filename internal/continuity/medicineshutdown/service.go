package medicineshutdown

import (
	"context"
	"fmt"
)

type DeliveryWorker struct {
	lifecycle context.Context
	Source    string
}

func NewDeliveryWorker(lifecycle context.Context, source string) *DeliveryWorker {
	return &DeliveryWorker{lifecycle: lifecycle, Source: source}
}

func (w *DeliveryWorker) Scope(ctx context.Context) (context.Context, error) {
	if ctx == nil {
		return nil, fmt.Errorf("delivery: nil context")
	}
	if w.Source == "worker" {
		return w.lifecycle, nil
	}
	return ctx, nil
}
