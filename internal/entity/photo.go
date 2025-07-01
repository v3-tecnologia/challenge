package entity

import (
	"time"

	"github.com/igorlopes88/desafio-v3/pkg/entity"
)

type Photo struct {
	ID         entity.ID `json:"id" gorm:"primaryKey"`
	User       entity.ID `json:"user" gorm:"not null"`
	MacAddress string    `json:"mac_address" gorm:"not null"`
	Image      string    `json:"image" gorm:"not null"`
	TimeStamp  time.Time `json:"timestamp" gorm:"not null"`
}

func NewPhoto(userId entity.ID, mac_address string, image string, time string) (*Photo, error) {
	newtime, err := SrtToTime(time)
	if err != nil {
		return nil, err
	}
	photo := &Photo{
		ID:         entity.NewID(),
		User:       userId,
		MacAddress: mac_address,
		Image:      image,
		TimeStamp:  newtime,
	}
	err = photo.Validate()
	if err != nil {
		return nil, err
	}
	// photo.GenerateImage()
	return photo, nil
}

func (p *Photo) GenerateImage() {
	// img, _, err := image.Decode(bytes.NewReader(p.Image))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// out, _ := os.Create("./img.jpg")
	// defer out.Close()

	// var opts jpeg.Options
	// opts.Quality = 1

	// err = jpeg.Encode(out, img, &opts)
	// if err != nil {
	// 	log.Println(err)
	// }
	// err := os.WriteFile("img.jpg", p.Image, 0644)
	// log.Println(err)
}

func (p *Photo) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.User.String() == "" {
		return ErrInvalidUser
	}
	if _, err := entity.ParseID(p.User.String()); err != nil {
		return ErrInvalidUser
	}
	if p.MacAddress == "" {
		return ErrMacIsRequired
	}
	return nil
}
