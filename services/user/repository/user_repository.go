package repository

import (
	"github.com/loak155/microservices/services/user/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *domain.User) error
	GetUser(user *domain.User, id int) error
	GetUserByEmail(user *domain.User, email string) error
	ListUsers(users *[]domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *domain.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUser(user *domain.User, id int) error {
	if err := ur.db.First(user, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(user *domain.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) ListUsers(users *[]domain.User) error {
	if err := ur.db.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *domain.User) error {
	if err := ur.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(id int) error {
	if err := ur.db.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
