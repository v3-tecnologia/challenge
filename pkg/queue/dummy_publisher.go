package queue

import (
	"context"
	"fmt"

	"github.com/dryingcore/v3-challenge/internal/core/ports/queue"
)

type DummyPublisher struct{}

func (p *DummyPublisher) Publish(ctx context.Context, subject string, msg queue.Message) error {
	fmt.Printf("[QUEUE] Subject: %s | Message: %+v\n", subject, msg)
	return nil
}
