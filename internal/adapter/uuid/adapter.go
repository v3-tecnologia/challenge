package uuid

import "github.com/google/uuid"

type Adapter struct{}

func (u *Adapter) Value() string {
	return uuid.New().String()
}
