package service

import "github.com/yduman/go-hexarch/internal/domain/entity"

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
}
