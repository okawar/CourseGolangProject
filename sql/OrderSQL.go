package sql

import "golang_pr/entity"

func AddOrder(Order entity.Order) {
	database.Table("orders").Select("user_id", "amount", "item_id").Create(Order)
}

func ChangeOrder(id uint32, updts entity.Order) {
	database.Table("orders").Where("order_id = ?", id).Updates(updts)
}
func DeleteOrder(id uint32) {
	var Order entity.Order
	database.Table("orders").Delete(&Order, id)
}

func GetAllOrders(owner_id uint32) []*entity.Order {
	var Orders []*entity.Order
	database.Table("orders").Where("user_id = ?", owner_id).Find(&Orders)
	return Orders
}

func GetOrderById(id uint32) entity.Order {
	var Order entity.Order
	database.Table("orders").Where("order_id = ?", id).First(&Order)
	return Order
}
