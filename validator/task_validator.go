package validator

import (
	"go-react-todo-udemy/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(tdas model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("タイトルは必須です"),
			validation.RuneLength(1, 10).Error("タイトルは1文字以上10文字以下です"),
		),
	)
}
