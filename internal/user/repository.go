package user

import (
	"errors"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	Get(where string, args ...interface{}) (*entity.User, error)
	Create(user *entity.User) error
	FirstOrCreate(user *entity.User, where string, args ...interface{}) (*entity.User, error)
	Update(user *entity.User, where string, args ...interface{}) error
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

func (r *repository) Get(where string, args ...interface{}) (*entity.User, error) {
	user := entity.User{}
	err := r.db.Where(where, args...).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		r.logger.Error("user repo get err: ", log.Any("err", err))
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		r.logger.Error("user repo create err: ", log.Any("err", err))
		return err
	}
	return nil
}

func (r *repository) FirstOrCreate(user *entity.User, where string, args ...interface{}) (*entity.User, error) {
	if err := r.db.Where(where, args).FirstOrCreate(user).Error; err != nil {
		r.logger.Error("user repo FirstOrCreate err: ", log.Any("err", err))
		return nil, err
	}
	return user, nil
}

func (r *repository) Update(user *entity.User, where string, args ...interface{}) error {
	if err := r.db.Model(user).Where(where, args).Updates(user).Error; err != nil {
		r.logger.Error("user repo update err: ", log.Any("err", err))
		return err
	}
	return nil
}
