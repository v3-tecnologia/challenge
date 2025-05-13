package uuid

import "github.com/google/uuid"

type Adapter struct{}
type UUIDGenerator interface {
	Generate() string
}

func (u *Adapter) Generate() string {
	return uuid.New().String()
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
