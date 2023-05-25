package aki

import (
	"aki/model"
	"aki/repository"

	"gorm.io/gorm"
)

type Aki struct {
	repos repository.Repository
}

func New(db *gorm.DB) *Aki {
	return &Aki{
		repos: *repository.New(db),
	}
}

func (a *Aki) GetUserByUID(uid string) (*model.User, error) {
	return a.repos.User.UserByUID(uid)
}
