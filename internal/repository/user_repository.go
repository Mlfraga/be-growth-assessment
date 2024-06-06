package repository

import (
	"go_api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *domain.User) error
    GetUsers() ([]domain.User, error)
    GetUserByID(id uint) (*domain.User, error)
    UpdateUser(user *domain.User) error
    DeleteUser(id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetUsers() ([]domain.User, error) {
    var users []domain.User
    return users, r.db.Find(&users).Error
}

func (r *userRepository) GetUserByID(id uint) (*domain.User, error) {
    var user domain.User
    return &user, r.db.Where("id = ?", id).First(&user).Error
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
