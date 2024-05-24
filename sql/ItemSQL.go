package sql

import (
	"golang_pr/entity"
)

func AddItem(item entity.Item) { //Экв. CreateItem в пакете service
	database.Table("items").Select("article", "name", "price", "creator", "description").Create(item)
}

func ChangeItem(id uint32, updts entity.Item) { //Экв. UpdateItem в пакете service
	database.Table("items").Where("item_id = ?", id).Updates(updts)
}

func DeleteItem(id uint32) { //Экв. DeleteItem в пакете service
	var item entity.Item
	database.Table("items").Delete(&item, id)
}

func GetItemById(id uint32) entity.Item { //Экв. GetItemById в пакете service
	var item entity.Item
	database.Table("items").Where("item_id = ?", id).First(&item)
	return item
}

func GetItemAll() []*entity.Item { //Экв. GetItemAll в пакете service
	var items []*entity.Item
	database.Table("items").Find(&items)
	return items
}
