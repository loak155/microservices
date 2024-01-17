package usecase

import (
	"github.com/loak155/microservices/services/user/domain"
	"github.com/loak155/microservices/services/user/repository"
	"github.com/loak155/microservices/services/user/validator"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(id int) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	ListUsers() ([]domain.User, error)
	UpdateUser(user domain.User) (bool, error)
	DeleteUser(id int) (bool, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) CreateUser(user domain.User) (domain.User, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return domain.User{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return domain.User{}, err
	}
	newUser := domain.User{Username: user.Username, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

func (uu *userUsecase) GetUser(id int) (domain.User, error) {
	storedUser := domain.User{}
	if err := uu.ur.GetUser(&storedUser, id); err != nil {
		return domain.User{}, err
	}
	return storedUser, nil
}

func (uu *userUsecase) GetUserByEmail(email string) (domain.User, error) {
	storedUser := domain.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, email); err != nil {
		return domain.User{}, err
	}
	return storedUser, nil
}

func (uu *userUsecase) ListUsers() ([]domain.User, error) {
	storedUsers := []domain.User{}
	if err := uu.ur.ListUsers(&storedUsers); err != nil {
		return []domain.User{}, err
	}
	return storedUsers, nil
}

func (uu *userUsecase) UpdateUser(user domain.User) (bool, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return false, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return false, err
	}
	updatedUser := domain.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.UpdateUser(&updatedUser); err != nil {
		return false, err
	}
	return true, nil
}

func (uu *userUsecase) DeleteUser(id int) (bool, error) {
	if err := uu.ur.DeleteUser(id); err != nil {
		return false, err
	}
	return true, nil
}
