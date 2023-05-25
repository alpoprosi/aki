package repository

import (
	"aki/model"

	"gorm.io/gorm"
)

type Repository struct {
	User UserRepository
}

type UserRepository interface {
	Save(user *model.User) error
	SaveUID(id uint, uid string) error

	User(id int64) (*model.User, error)
	Users(query, orderBy string, limit, offset int) ([]model.User, int64, error)
	UserByUID(uid string) (*model.User, error)

	WithTrx(*gorm.DB) *userRepository
}
