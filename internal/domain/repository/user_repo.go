package repository

import "github.com/yduman/go-hexarch/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id string) (*entity.User, error)
}
