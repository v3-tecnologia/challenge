package queue

import core "github.com/dryingcore/v3-challenge/internal/core/ports/queue"

func NewFakePublisher() core.Publisher {
	return &DummyPublisher{}
}

func NewRealPublisher(url string) (core.Publisher, error) {
	return NewNatsPublisher(url)
}
