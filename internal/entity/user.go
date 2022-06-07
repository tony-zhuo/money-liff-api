package entity

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

var validate *validator.Validate

type User struct {
	Id        int            `json:"-" validate:"-" gorm:"primaryKey"`
	LineId    string         `json:"line_id" validate:"required,alphanum" binding:"required"`
	Name      string         `json:"name" validate:"required,alphanum" binding:"required"`
	AvatarUrl string         `json:"avatar_url" validate:"-"`
	CreatedAt time.Time      `json:"-" validate:"-"`
	UpdatedAt time.Time      `json:"-" validate:"-"`
	DeletedAt gorm.DeletedAt `json:"-" validate:"-"`
}

func (u *User) GetLineId() string {
	return u.LineId
}

func (u *User) Validate() error {
	validate = validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
