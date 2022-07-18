package user

import (
	"errors"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	Get(lineId string) *entity.User
	Create(user *entity.User) error
	FirstOrCreate(user *entity.User) error
	Update(user *entity.User) error
}

type repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
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
	if result := r.db.Debug().Create(user); result.Error != nil {
		r.logger.Error("user repo create err: ", log.Any("err", result.Error))
		return result.Error
	}
	return nil
}

func (r *repository) FirstOrCreate(user *entity.User) error {
	if result := r.db.Where(entity.User{LineId: user.LineId}).FirstOrCreate(user); result.Error != nil {
		r.logger.Error("user repo FirstOrCreate err: ", log.Any("err", result.Error))
		return result.Error
	}
	return nil
}

func (r *repository) Update(user *entity.User) error {
	if result := r.db.Model(user).Where("line_id = ?", user.LineId).Updates(user); result.Error != nil {
		r.logger.Error("user repo update err: ", log.Any("err", result.Error))
		return result.Error
	}
	return nil
}
