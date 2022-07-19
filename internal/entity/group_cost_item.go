package entity

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type GroupCostItem struct {
	Id           int                        `json:"-" gorm:"primaryKey"`
	GroupId      int                        `json:"group_id"`
	UUID         string                     `json:"uuid" `
	Name         string                     `json:"name"`
	TotalAmount  int                        `json:"total_amount"`
	PayerId      int                        `json:"-"`
	CreatorId    int                        `json:"-"`
	PayAt        time.Time                  `json:"pay_at"`
	Remark       string                     `json:"remark"`
	Participants []GroupCostItemParticipant `gorm:"foreignKey:group_cost_item_id;"`
	CreatedAt    time.Time                  `json:"-" validate:"-"`
	UpdatedAt    time.Time                  `json:"-" validate:"-"`
	DeletedAt    gorm.DeletedAt             `json:"-" validate:"-"`
	Payer        User
}

type GroupCostItemRequestArg struct {
	Name         string                               `json:"name" validate:"required"`
	TotalAmount  int                                  `json:"total_amount" validate:"required"`
	PayerId      int                                  `json:"payer_id" validate:"required"`
	PayAt        time.Time                            `json:"pay_at" form:"pay_at" time_format:"unixNano"`
	Remark       string                               `json:"remark"`
	Participants []GroupCostItemParticipantRequestArg `json:"participants"`
}

func (g *GroupCostItemRequestArg) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}
	return nil
}

type GroupCostItemResponse struct {
	Name        string        `json:"name"`
	TotalAmount int           `json:"total_amount"`
	PayAt       *string       `json:"pay_at"`
	Payer       PayerResponse `json:"payer"`
}
