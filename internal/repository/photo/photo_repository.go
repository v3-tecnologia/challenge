package photo

import "github.com/iamrosada0/v3/internal/domain"

type PhotoRepository interface {
	Create(d *domain.Photo) (*domain.Photo, error)
}
