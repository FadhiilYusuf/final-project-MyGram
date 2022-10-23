package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Mesagge is required"`
	UserID  int    ` json:"user_id"`
	PhotoID int    ` json:"photo_id" form:"photo_id" valid:"required~Photo ID is required"`
	User    *User
	Photo   *Photo
}

type GetAllCommentsResponse struct {
	GormModel
	Message string ` json:"message" `
	PhotoID int    `json:"photo_id" `
	UserID  int    `gorm:"foreignKey" json:"user_id"`
	Photo   struct {
		ID       int    `json:"photo_id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string ` json:"photo_url"`
	}
	User struct {
		Email    string `json:"email"`
		UserName string `json:"user_name"`
	}
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
