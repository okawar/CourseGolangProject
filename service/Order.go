package service

import (
	"golang_pr/entity"
	"golang_pr/sql"
	"log"
	"strconv"
)

func init() {
	CRUDS["OrderCreate"] = CreateOrder
	CRUDS["OrderUpdate"] = UpdateOrder
	CRUDS["OrderDelete"] = DeleteOrder
	CRUDS["Orders"] = GetOrderAll
	CRUDS["Order"] = GetOrderById
}

func CreateOrder(ctx *Context) { //HTTP.POST
	var Order entity.Order
	Order, err := GetData[entity.Order](ctx)
	if err != nil {
		log.Println("Error decoding json for create order.", err)
		ctx.sendError(500, "Error decoding json for create order.")
	}
	sql.AddOrder(Order)
}

func UpdateOrder(ctx *Context) { //HTTP.PUT
	var Order entity.Order
	Order, err := GetData[entity.Order](ctx)
	if err != nil {
		log.Println("Error of decoding json for updating order.\n", err)
		ctx.sendError(500, "Error ofjson decoding.")
	}
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for update order.\n", err)
		ctx.sendError(400, "Inavlid url for updating order.")
	}
	sql.ChangeOrder(uint32(id), Order)
}

func DeleteOrder(ctx *Context) { //HTTP.DELETE
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for delete order.\n", err)
		ctx.sendError(400, "Inavlid url for deleting order.")
	}
	sql.DeleteOrder(uint32(id))
}

func GetOrderById(ctx *Context) { //HTTP.GET
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for getting order.\n", err)
		ctx.sendError(400, "Inavlid url for getting order.")
	}
	ctx.sendAnswer(sql.GetOrderById(uint32(id)))
}

func GetOrderAll(ctx *Context) { //HTTP.GET
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["owner_id"][0], 10, 32)
	if err != nil {
		log.Println("Error of owner_id reading for getting orders.\n", err)
		ctx.sendError(400, "Inavlid url for getting orders.")
	}
	ctx.sendAnswer(sql.GetAllOrders(uint32(id)))
}
