package entities

type PhotoEntity struct {
	BaseEntity
	FilePath string `json:"file_path" gorm:"column:file_path"`
}

func (PhotoEntity) TableName() string {
	return "photo"
}
