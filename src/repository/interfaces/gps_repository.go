package interfaces

import "v3-backend-challenge/model"

type GpsRepository interface {
	Save(photo *model.GPS) error
}
