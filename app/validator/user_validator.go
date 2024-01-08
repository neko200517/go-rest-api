package validator

import (
	"go-echo/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// バリデーションの実装
func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email, // Email
			validation.Required.Error("email is required"),            // 必須
			validation.RuneLength(1, 30).Error("limited max 30 char"), // 1～30文字
			is.Email.Error("is not valid email format"),               // Emailフォーマットの検証
		),
		validation.Field(
			&user.Password, // Password
			validation.Required.Error("password is required"),               // 必須
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"), // 6～30文字
		),
	)
}
