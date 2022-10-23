package models

import (
	"errors"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName     string        `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~User name is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)" `
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	if u.Age <= 8 {
		err = errors.New("age must be higher than 8")

		return err
	}

	hashedPass := helpers.HashPass(u.Password)

	u.Password = hashedPass
	return nil
}
