package entities

type GyroscopeEntity struct {
	BaseEntity
	XValue float64 `json:"x_value" gorm:"column:x_value;not null"`
	YValue float64 `json:"y_value" gorm:"column:y_value;not null"`
	ZValue float64 `json:"z_value" gorm:"column:z_value;not null"`
}

func (GyroscopeEntity) TableName() string {
	return "gyroscope"
}
