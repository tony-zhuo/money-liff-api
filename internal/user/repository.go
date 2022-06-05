package user

import (
	"errors"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/database"
	"gorm.io/gorm"
)

type Repository interface {
	Get(lineId string) *entity.User
	Create(user *entity.User) error
	Update(user *entity.User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		db: database.Connection(),
	}
}

func (r *repository) Get(lineId string) *entity.User {
	user := entity.User{}
	result := r.db.Where("line_id = ?", lineId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (r *repository) Create(user *entity.User) error {
	r.db.Create(user)
	return nil
}

func (r *repository) Update(user *entity.User) error {
	r.db.Model(user).Where("line_id = ?", user.LineId).Updates(user)
	return nil
}
