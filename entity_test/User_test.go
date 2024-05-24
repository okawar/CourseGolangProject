package entity_test

import (
	"golang_pr/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	// Создание экземпляра User с недопустимыми значениями
	invalidUser := entity.User{
		UserId:   0,               // Недопустимое значение для UserId
		FIO:      "",              // Пустая строка для FIO
		Email:    "invalid_email", // Некорректный формат email
		Password: "short",         // Слишком короткий пароль
		Login:    "user123",       // Недопустимый логин
		Status:   5,               // Недопустимое значение для Status
	}

	// Валидация
	errs := invalidUser.Validate()

	// Проверка ошибок валидации
	assert.NotNil(t, errs, "Errors should not be nil for invalid user")

	// Проверка наличия ожидаемых ошибок
	assert.Contains(t, errs, "UserId must be greater than 0", "UserId validation error should be present")
	assert.Contains(t, errs, "FIO cannot be empty", "FIO validation error should be present")
	assert.Contains(t, errs, "Email is not valid", "Email validation error should be present")
	assert.Contains(t, errs, "DateBirth cannot be zero", "DateBirth validation error should be present")
	assert.Contains(t, errs, "Password length must be at least 6 characters", "Password validation error should be present")
	assert.Contains(t, errs, "Login is invalid", "Login validation error should be present")
	assert.Contains(t, errs, "Status must be either 0 or 1", "Status validation error should be present")
}

func TestUserValidationSuccess(t *testing.T) {
	// Создание экземпляра User с допустимыми значениями
	validUser := entity.User{
		UserId:   1,
		FIO:      "John Doe",
		Email:    "john@example.com",
		Password: "strongpassword",
		Login:    "user1234",
		Status:   1,
	}

	// Валидация
	errs := validUser.Validate()

	// Проверка, что ошибок нет
	assert.Nil(t, errs, "No errors should be present for valid user")
}
