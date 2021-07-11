package services

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/repository"
)

type UserService interface {
	GetOwnProjects(email string) (*entity.User, error)
	CreateNewUser(user *entity.User) (*entity.User, error)
}

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (us *userService) GetOwnProjects(id string) (*entity.User, error) {
	return us.UserRepo.GetOwnProjects(id)
}

func (us *userService) CreateNewUser(user *entity.User) (*entity.User, error) {
	return us.UserRepo.CreateNewUser(user)
}
