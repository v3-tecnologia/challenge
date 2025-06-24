package queue

import (
	"context"
	"encoding/json"

	core "github.com/dryingcore/v3-challenge/internal/core/ports/queue"
	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewNatsPublisher(url string) (core.Publisher, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsPublisher{conn: nc}, nil
}

func (p *NatsPublisher) Publish(ctx context.Context, subject string, msg core.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.conn.Publish(subject, data)
}
