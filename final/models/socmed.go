package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Socmed struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name" form:"name" valid:"required~nama tidak boleh kosong"`
	Socmed_URL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~URL tidak boleh kosong"`
	User_ID    uint   `json:"user_id"`
	User       *User
}

func (s *Socmed) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
