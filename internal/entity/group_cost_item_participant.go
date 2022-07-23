package entity

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/validate_err_msg"
	"github.com/go-playground/validator/v10"
)

type GroupCostItemParticipant struct {
	Id              int `json:"-" gorm:"primaryKey"`
	GroupCostItemId int `json:"group_cost_item_id"`
	UserId          int `json:"-"`
	Amount          int `json:"amount"`
}

type GroupCostItemParticipantRequestArg struct {
	UserId int `json:"user_id" validate:"required,numeric"`
	Amount int `json:"amount" validate:"required,numeric"`
}

func (g *GroupCostItemParticipantRequestArg) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return validate_err_msg.Transfer(err)
	}
	return nil
}
