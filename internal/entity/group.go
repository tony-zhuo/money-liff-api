package entity

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/validate_err_msg"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	Id          int            `json:"-" validate:"-" gorm:"primaryKey"`
	UUID        string         `json:"uuid" validate:"-"`
	Name        string         `json:"name" validate:"required" binding:"required"`
	UserLimit   int            `json:"user_limit" validate:"required,numeric" binding:"required"`
	ImageUrl    string         `json:"image_url" validate:"-"`
	AdminUserId int            `json:"-" validate:"-" gorm:"column:admin_user_id"`
	CreatedAt   time.Time      `json:"-" validate:"-"`
	UpdatedAt   time.Time      `json:"-" validate:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" validate:"-"`
	Users       []*User        `json:"-" gorm:"many2many:user_groups;"`
	CostItem    []GroupCostItem
}

type GroupWithCostItemResponse struct {
	UUID      string                  `json:"uuid"`
	Name      string                  `json:"name"`
	UserLimit int                     `json:"user_limit"`
	ImageUrl  string                  `json:"image_url"`
	CostItem  []GroupCostItemResponse `json:"cost_item"`
}

type GroupResponse struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	IsAdmin  bool   `json:"is_admin"`
}

func (g *Group) GetUuid() string {
	return g.UUID
}

func (g *Group) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return validate_err_msg.Transfer(err)
	}
	return nil
}

type GroupParams struct {
	UUID string `json:"uuid" validate:"uuid"`
}

func (g *GroupParams) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return validate_err_msg.Transfer(err)
	}
	return nil
}
