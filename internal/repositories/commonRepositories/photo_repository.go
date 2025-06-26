package commonRepositories

import "v3-test/internal/models/commonModels"

type PhotoRepository interface {
	CreatePhoto(photoModel commonModels.PhotoModel) (commonModels.PhotoModel, error)
}
