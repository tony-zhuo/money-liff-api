package entity

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	Id          int            `json:"-" validate:"-" gorm:"primaryKey"`
	UUID        string         `json:"uuid" validate:"-"`
	Name        string         `json:"name" validate:"required,ascii" binding:"required"`
	UserLimit   int            `json:"user_limit" validate:"required,numeric" binding:"required"`
	ImageUrl    string         `json:"image_url" validate:"-"`
	AdminUserId int            `json:"-" validate:"-" gorm:"column:admin_user_id"`
	CreatedAt   time.Time      `json:"-" validate:"-"`
	UpdatedAt   time.Time      `json:"-" validate:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" validate:"-"`
	Users       []*User        `json:"-" gorm:"many2many:user_groups;"`
}

func (g *Group) GetUuid() string {
	return g.UUID
}

func (g *Group) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}
	return nil
}

type GroupParams struct {
	UUID string `json:"uuid" validate:"uuid"`
}

func (g *GroupParams) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}
	return nil
}
