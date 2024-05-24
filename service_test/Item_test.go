package service_test

import (
	"golang_pr/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateItem(t *testing.T) {
	itemJSON := `{"id":1,"name":"Item 1","price":10.5}`
	req, err := http.NewRequest("POST", "/api/item/new", strings.NewReader(itemJSON))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	ctx := &service.Context{
		Response: rec,
		Request:  req,
	}

	service.CreateItem(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rec.Code)
	}
}

func TestUpdateItem(t *testing.T) {
	itemJSON := `{"id":1,"name":"Updated Item","price":15.75}`
	req, err := http.NewRequest("PUT", "/api/item/update?id=1", strings.NewReader(itemJSON))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	ctx := &service.Context{
		Response: rec,
		Request:  req,
	}

	service.UpdateItem(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/item/delete?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	ctx := &service.Context{
		Response: rec,
		Request:  req,
	}

	service.DeleteItem(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rec.Code)
	}
}

func TestGetItemById(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/item/get?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	ctx := &service.Context{
		Response: rec,
		Request:  req,
	}

	service.GetItemById(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rec.Code)
	}
}

func TestGetItemAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/item/all", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	ctx := &service.Context{
		Response: rec,
		Request:  req,
	}

	service.GetItemAll(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rec.Code)
	}
}
