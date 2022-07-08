package entity

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type GroupCostItem struct {
	Id          int       `json:"-" validate:"-" gorm:"primaryKey"`
	GroupId     int       `json:"group_id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	TotalAmount int       `json:"total_amount"`
	PayerId     int       `json:"-"`
	CreatorId   int       `json:"-"`
	PayAt       time.Time `json:"pay_at"`
	Remark      string    `json:"remark"`
}

func (g *GroupCostItem) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}
	return nil
}
