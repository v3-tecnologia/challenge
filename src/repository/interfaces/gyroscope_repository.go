package interfaces

import "v3-backend-challenge/model"

type GyroscopeRepository interface {
	Save(gyroscope *model.Gyroscope) error
}
