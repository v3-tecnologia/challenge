package common

import "time"

type PhotoModel struct {
	ID  string `json:"id" bson:"_id"`
	Url string `json:"url" bson:"url"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
