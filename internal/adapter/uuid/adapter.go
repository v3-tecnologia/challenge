package uuid

import "github.com/google/uuid"

type Adapter struct{}
type AdapterInterface interface {
	Value() string
}

func (u *Adapter) Value() string {
	return uuid.New().String()
}
