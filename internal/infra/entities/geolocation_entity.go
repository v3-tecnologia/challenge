package entities

type GeolocationEntity struct {
	BaseEntity
	Latitude  float64 `json:"latitude" gorm:"column:latitude;not null"`
	Longitude float64 `json:"longitude" gorm:"column:longitude;not null"`
}

func (GeolocationEntity) TableName() string {
	return "geolocation"
}
