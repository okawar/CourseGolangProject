package service_test

import (
	"golang_pr/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Подготовка тестовых данных
	userJSON := `{"user_id":1,"fio":"John Doe","email":"john@example.com","password":"password123","login":"johndoe","status":1}`
	req, err := http.NewRequest("POST", "/api/user/new", strings.NewReader(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	// Создание контекста для тестирования
	context := &service.Context{Response: rec, Request: req}

	// Вызов функции CreateUser
	service.CreateUser(context)

	// Проверка статуса ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидаемый статус %d; получен %d", http.StatusOK, rec.Code)
	}

	// Проверка тела ответа
	expectedResponse := `{"mes":"granted"}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Ожидаемое тело ответа %s; получено %s", expectedResponse, rec.Body.String())
	}
}

func TestUpdateUser(t *testing.T) {
	// Подготовка тестовых данных
	userJSON := `{"user_id":1,"fio":"John Doe","email":"john@example.com","password":"newpassword123","login":"johndoe","status":1}`
	req, err := http.NewRequest("PUT", "/api/user/update", strings.NewReader(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	// Создание контекста для тестирования
	context := &service.Context{Response: rec, Request: req}

	// Вызов функции UpdateUser
	service.UpdateUser(context)

	// Проверка статуса ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидаемый статус %d; получен %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	// Подготовка тестовых данных
	req, err := http.NewRequest("DELETE", "/api/user/delete?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	// Создание контекста для тестирования
	context := &service.Context{Response: rec, Request: req}

	// Вызов функции DeleteUser
	service.DeleteUser(context)

	// Проверка статуса ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидаемый статус %d; получен %d", http.StatusOK, rec.Code)
	}
}

func TestGetUserById(t *testing.T) {
	// Подготовка тестовых данных
	req, err := http.NewRequest("GET", "/api/user/get?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	// Создание контекста для тестирования
	context := &service.Context{Response: rec, Request: req}

	// Вызов функции GetUserById
	service.GetUserById(context)

	// Проверка статуса ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидаемый статус %d; получен %d", http.StatusOK, rec.Code)
	}
}

func TestGetUserAll(t *testing.T) {
	// Подготовка тестовых данных
	req, err := http.NewRequest("GET", "/api/user/all", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	// Создание контекста для тестирования
	context := &service.Context{Response: rec, Request: req}

	// Вызов функции GetUserAll
	service.GetUserAll(context)

	// Проверка статуса ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Ожидаемый статус %d; получен %d", http.StatusOK, rec.Code)
	}
}
