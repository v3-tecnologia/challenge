package queue

import (
	"context"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Publisher interface {
	Publish(ctx context.Context, subject string, msg Message) error
}
