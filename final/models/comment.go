package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	User_ID  uint   `json:"user_id"`
	Photo_ID uint   `json:"photo_id"`
	Message  string `gorm:"not null" json:"message" form:"message" valid:"required~input message"`
	User     *User  `gorm:"foreignKey:User_ID"`
	Photo    *Photo `gorm:"foreignKey:Photo_ID"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
