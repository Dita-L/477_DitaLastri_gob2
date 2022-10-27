package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	Caption   string `json:"caption" form:"caption"`
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~please input photo title"`
	Photo_URL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~please upload"`
	User_ID   uint   `json:"user_id"`
	gorm.Model
	User User `gorm:"foreignKey:User_ID"`
	//Comment []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
