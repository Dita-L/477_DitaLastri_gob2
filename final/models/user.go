package models

import (
	"final/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `gorm:"not null;unique_index" json:"user_name" form:"user_name" valid:"required~please input username"`
	Email    string `gorm:"not null;unique_index" json:"email" form:"email" valid:"email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~please input password,minstringlength(6)~minimal 6 karakter"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"range(8|100)"`
	gorm.Model
}

// Password di hash dengan bcrypt sebelum disimpan ke database
func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

// untuk update, perlu validasi lagi, disertai autentikasi
// func (u *User) BeforeUpdate(ctx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }
