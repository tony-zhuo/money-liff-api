package entity

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/validate_err_msg"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

var validate *validator.Validate

type User struct {
	Id        int            `json:"-" validate:"-" gorm:"primaryKey"`
	LineId    string         `json:"line_id" validate:"required,alphanum" binding:"required"`
	Name      string         `json:"name" validate:"required" binding:"required"`
	AvatarUrl string         `json:"avatar_url" validate:"-"`
	CreatedAt time.Time      `json:"-" validate:"-"`
	UpdatedAt time.Time      `json:"-" validate:"-"`
	DeletedAt gorm.DeletedAt `json:"-" validate:"-"`
	Groups    []Group        `json:"-" gorm:"many2many:user_groups;"`
}

func (u *User) GetLineId() string {
	return u.LineId
}

func (u *User) Validate() error {
	validate = validator.New()
	if err := validate.Struct(u); err != nil {
		return validate_err_msg.Transfer(err)
	}
	return nil
}

type PayerResponse struct {
	Name      string  `json:"name"`
	AvatarUrl *string `json:"avatar_url"`
}
