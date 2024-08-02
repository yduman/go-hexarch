package service

import (
	"github.com/yduman/go-hexarch/internal/domain/entity"
	"github.com/yduman/go-hexarch/internal/domain/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (svc *userService) CreateUser(user *entity.User) error {
	return svc.repo.Save(user)
}

func (svc *userService) GetUserByID(id string) (*entity.User, error) {
	return svc.repo.FindByID(id)
}
