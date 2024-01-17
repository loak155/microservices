package validator

import (
	"github.com/loak155/microservices/services/user/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user domain.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user domain.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Username,
			validation.Required.Error("username is required"),
			validation.RuneLength(6, 20).Error("limited min 6 max 20 char"),
		),
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limited min 8 max 30 char"),
		),
	)
}
