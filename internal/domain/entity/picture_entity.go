package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Picture struct {
	BaseEntity
	PictureURL       string  `json:"picture_url" gorm:"type:varchar(255);not null"`
	PictureFormat    string  `json:"picture_format" gorm:"type:varchar(255);not null"`
	RecognizedFace   bool    `json:"recognized_face" gorm:"type:boolean;default:false"`
	RekognitionScore float64 `json:"rekognition_score,omitempty" gorm:"type:double precision"`
}

func BuildPictures(deviceId string, createdAt time.Time, pictureUrl, pictureFormat string, recognizedFace bool, rekognitionScore float64) *Picture {
	return &Picture{
		BaseEntity: BaseEntity{
			DeviceID:   deviceId,
			CreatedAt:  createdAt,
			ReceivedAt: time.Now(),
		},
		PictureURL:       pictureUrl,
		PictureFormat:    pictureFormat,
		RecognizedFace:   recognizedFace,
		RekognitionScore: rekognitionScore,
	}
}

func (g *Picture) Validate() error {
	if g.ID == uuid.Nil {
		return errors.New("o ID da telemetria do giroscópio não pode ser nulo")
	}

	if g.DeviceID == "" {
		return errors.New("o ID do dispositivo não pode estar vazio")
	}

	if g.CreatedAt.IsZero() {
		return errors.New("a data de criação não pode ser zero")
	}

	if g.ReceivedAt.IsZero() {
		return errors.New("a data de recebimento não pode ser zero")
	}

	if g.PictureURL == "" {
		return errors.New("a url da imagem nao pode ser vazia")
	}

	return nil
}
