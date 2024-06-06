package service

import (
	"go_api/internal/domain"
	"go_api/internal/repository"
)

type UserService interface {
    CreateUser(user *domain.User) error
    GetUsers() ([]domain.User, error)
    GetUserByID(id uint) (*domain.User, error)
    UpdateUser(user *domain.User) error
    DeleteUser(id uint) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo}
}

func (s *userService) CreateUser(user *domain.User) error {
    return s.repo.CreateUser(user)
}

func (s *userService) GetUsers() ([]domain.User, error) {
    return s.repo.GetUsers()
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
    return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *domain.User) error {
    var userDB, err = s.repo.GetUserByID(user.ID)

    if err != nil {
        return err
    }

    userDB.Name = user.Name
    userDB.Email = user.Email
    userDB.Phone = user.Phone
    userDB.Document = user.Document

    return s.repo.UpdateUser(userDB)
}

func (s *userService) DeleteUser(id uint) error {
    return s.repo.DeleteUser(id)
}
