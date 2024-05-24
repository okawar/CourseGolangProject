package entity

import (
	"errors"
	"regexp"
)

type User struct {
	UserId   uint32 `json: "user_id"`
	FIO      string `json: "fio"`
	Email    string `json: "email"`
	Password string `json: "-"`
	Login    string `json: "login"`
	Status   uint32 `json: "status"`
}

// Validate проверяет корректность полей User.
// Возвращает срез ошибок, если какое-либо из полей недопустимо.
func (user *User) Validate() []error {
	var errs []error

	if user.FIO == "" {
		errs = append(errs, errors.New("FIO CAN'T BE EMPTY"))
	}

	if !isValidEmail(user.Email) {
		errs = append(errs, errors.New("INVALID EMAIL"))
	}

	if len(user.Password) < 6 {
		errs = append(errs, errors.New("PASSWORD SHORTER THAN 6 CHARS"))
	}

	if !isValidLogin(user.Login) {
		errs = append(errs, errors.New("LOGIN HAS INVALID CHARS"))
	}

	return errs
}

// isValidEmail проверяет, является ли переданный email допустимым.
func isValidEmail(email string) bool {
	// Простая проверка формата email
	// Здесь можно добавить более сложную проверку
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// isValidLogin проверяет, является ли переданный логин допустимым.
func isValidLogin(login string) bool {
	// Пример простой проверки на допустимые символы в логине
	loginRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return loginRegex.MatchString(login)
}
