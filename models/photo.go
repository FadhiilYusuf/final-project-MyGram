package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string ` gorm:"not null" json:"title" form:"title" valid:"required~Photo title is required"`
	Caption  string ` json:"caption" form:"caption" valid:"optional" `
	PhotoUrl string ` gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   int    ` json:"user_id"`
	User     *User
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
}

type GetAllPhotosResponse struct {
	GormModel
	Title    string ` json:"title" `
	Caption  string ` json:"caption" `
	PhotoUrl string ` json:"photo_url" `
	UserID   int    ` json:"user_id"`
	User     struct {
		Email    string `json:"email"`
		UserName string `json:"user_name"`
	}
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
