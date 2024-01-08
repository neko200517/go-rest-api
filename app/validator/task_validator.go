package validator

import (
	"go-echo/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// バリデーションの実装
func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title, // 対象フィールド
			validation.Required.Error("title is required"),            // 入力必須
			validation.RuneLength(1, 10).Error("limited max 10 char"), // 1～10文字以内
		),
	)
}
