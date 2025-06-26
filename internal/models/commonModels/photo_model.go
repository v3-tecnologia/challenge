package commonModels

import (
	"time"
	"v3-test/internal/enums"
)

type PhotoModel struct {
	Url    string            `json:"url" bson:"url"`
	Entity enums.PhotoEntity `bson:"entity"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
