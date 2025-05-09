package interfaces

import "v3-backend-challenge/src/model"

type PhotoRepository interface {
	Save(photo *model.Photo) error
}
