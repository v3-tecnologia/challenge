package interfaces

import "v3-backend-challenge/model"

type PhotoRepository interface {
	Save(photo *model.Photo) error
}
