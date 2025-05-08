package uuid

import "github.com/google/uuid"

type Adapter struct{}
type AdapterInterface interface {
	Value() string
}

func (u *Adapter) Generate() string {
	return uuid.New().String()
}
func NewUUIDAdapter() *Adapter {
	return &Adapter{}
}
