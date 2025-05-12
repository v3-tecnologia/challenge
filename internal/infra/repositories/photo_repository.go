package repositories

import (
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
)

type PhotoRepository interface {
	Create(photo entities.PhotoEntity) (entities.PhotoEntity, error)
	GetAll() ([]entities.PhotoEntity, error)
	ListByMacAddress(macAddress string) ([]entities.PhotoEntity, error)
}

type photoRepository struct {
	db *database.Database
}

func NewPhotoRepository(db *database.Database) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (r *photoRepository) Create(photo entities.PhotoEntity) (entities.PhotoEntity, error) {
	if err := r.db.DB.Create(&photo).Error; err != nil {
		return entities.PhotoEntity{}, err
	}
	return photo, nil
}

func (r *photoRepository) GetAll() ([]entities.PhotoEntity, error) {
	var photos []entities.PhotoEntity
	if err := r.db.DB.Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (r *photoRepository) ListByMacAddress(macAddress string) ([]entities.PhotoEntity, error) {
	var photos []entities.PhotoEntity
	if err := r.db.DB.Where("mac_address = ?", macAddress).Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}
