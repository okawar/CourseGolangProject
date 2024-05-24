package entity

import (
	"encoding/json"
	"golang_pr/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemJSONSerialization(t *testing.T) {
	item := entity.Item{
		ItemId:      1,
		Article:     12345,
		Name:        "TestItem",
		Price:       10.99,
		Creator:     "TestCreator",
		Description: "This is a test item.",
	}

	// Сериализация в JSON
	jsonData, err := json.Marshal(item)
	assert.NoError(t, err, "Error should be nil during JSON marshaling")

	// Десериализация из JSON
	var newItem entity.Item
	err = json.Unmarshal(jsonData, &newItem)
	assert.NoError(t, err, "Error should be nil during JSON unmarshaling")

	// Проверка, что значения соответствуют ожидаемым
	assert.Equal(t, item.ItemId, newItem.ItemId, "Item IDs should match")
	assert.Equal(t, item.Article, newItem.Article, "Article numbers should match")
	assert.Equal(t, item.Name, newItem.Name, "Names should match")
	assert.Equal(t, item.Price, newItem.Price, "Prices should match")
	assert.Equal(t, item.Creator, newItem.Creator, "Creators should match")
	assert.Equal(t, item.Description, newItem.Description, "Descriptions should match")
}

func TestItemValidation(t *testing.T) {
	// Создание экземпляра Item с недопустимыми значениями
	invalidItem := entity.Item{
		ItemId:      0, // Недопустимое значение для ItemId
		Article:     12345,
		Name:        "",     // Пустая строка для Name
		Price:       -10.99, // Отрицательное значение для Price
		Creator:     "TestUser",
		Description: "This is a test item.",
	}

	// Валидация
	errs := invalidItem.Validate()

	// Проверка ошибок валидации
	assert.NotNil(t, errs, "Errors should not be nil for invalid item")

	// Проверка наличия ожидаемых ошибок
	assert.Contains(t, errs, "ItemId must be greater than 0", "ItemId validation error should be present")
	assert.Contains(t, errs, "Name cannot be empty", "Name validation error should be present")
	assert.Contains(t, errs, "Price must be greater than 0", "Price validation error should be present")
}

func TestItemValidationSuccess(t *testing.T) {
	// Создание экземпляра Item с допустимыми значениями
	validItem := entity.Item{
		ItemId:      1,
		Article:     12345,
		Name:        "TestItem",
		Price:       10.99,
		Creator:     "TestUser",
		Description: "This is a test item.",
	}

	// Валидация
	errs := validItem.Validate()

	// Проверка, что ошибок нет
	assert.Nil(t, errs, "No errors should be present for valid item")
}
