package interfaces

import "v3-backend-challenge/src/model"

type GyroscopeRepository interface {
	Save(gyroscope *model.Gyroscope) error
}
