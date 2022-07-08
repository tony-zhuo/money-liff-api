package entity

import (
	"github.com/go-playground/validator/v10"
)

type GroupCostItemParticipant struct {
	Id              int `json:"-" validate:"-" gorm:"primaryKey"`
	GroupId         int `json:"group_id"`
	GroupCostItemId int `json:"group_cost_item_id"`
	UserId          int `json:"-"`
	Amount          int `json:"amount"`
}

func (g *GroupCostItemParticipant) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}
	return nil
}
