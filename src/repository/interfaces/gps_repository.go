package interfaces

import "v3-backend-challenge/src/model"

type GpsRepository interface {
	Save(photo *model.GPS) error
}
