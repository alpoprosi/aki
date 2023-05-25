package repository

import (
	"aki/model"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

var _ UserRepository = (*userRepository)(nil)

func New(db *gorm.DB) *Repository {
	return &Repository{
		User: &userRepository{DB: db},
	}
}

func (u *userRepository) Save(user *model.User) error {
	err := u.DB.Model(user).Create(user).Error

	return err
}

func (u *userRepository) SaveUID(id uint, uid string) error {
	err := u.DB.Model(&model.User{}).
		Where("id = ?", id).
		Update("uid", uid).
		Error

	return err
}

func (u *userRepository) User(id int64) (*model.User, error) {
	var result *model.User
	err := u.DB.Model(result).
		Select("*").
		Where("id = ?", id).
		Find(result).
		Error

	if isNotFound(err) {
		return nil, nil
	}

	return result, err
}

func (u *userRepository) UserByUID(uid string) (*model.User, error) {
	var result *model.User
	err := u.DB.Model(result).
		Select("*").
		Where("uid = ?", uid).
		Find(result).
		Error

	if isNotFound(err) {
		return nil, nil
	}

	return result, err
}

func (u *userRepository) Users(
	query, orderBy string,
	limit, offset int,
) ([]model.User, int64, error) {
	count := int64(0)
	q := u.DB.Model(&model.User{}).Select("*")

	if query != "" {
		q = q.Where("value ilike ?", "%"+query+"%")
	}

	if orderBy != "" {
		q = q.Order(orderBy)
	}

	var result []model.User
	err := q.Count(&count).
		Limit(limit).
		Offset(offset).
		Find(&result).Error

	if isNotFound(err) {
		return nil, 0, nil
	}

	return result, count, err
}

func (u *userRepository) WithTrx(tx *gorm.DB) *userRepository {
	if tx == nil {
		log.Print("tx database not found")

		return u
	}

	u.DB = tx

	return u
}
