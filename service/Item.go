package service

import (
	"golang_pr/entity"
	"golang_pr/sql"
	"log"
	"strconv"
)

func init() {
	CRUDS["ItemCreate"] = CreateItem
	CRUDS["ItemUpdate"] = UpdateItem
	CRUDS["ItemDelete"] = DeleteItem
	CRUDS["Items"] = GetItemAll
	CRUDS["Item"] = GetItemById
}

func CreateItem(ctx *Context) { //HTTP.POST
	var item entity.Item
	item, err := GetData[entity.Item](ctx)
	if err != nil {
		log.Println("Error decoding json for create item.", err)
		ctx.sendError(500, "Error decoding json for create item.")
	}
	sql.AddItem(item)
}

func UpdateItem(ctx *Context) { //HTTP.PUT
	var item entity.Item
	item, err := GetData[entity.Item](ctx)
	if err != nil {
		log.Println("Error of decoding json for updating item.\n", err)
		ctx.sendError(500, "Error ofjson decoding.")
	}
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for update item.\n", err)
		ctx.sendError(400, "Inavlid url for updating item.")
	}
	sql.ChangeItem(uint32(id), item)
}

func DeleteItem(ctx *Context) { //HTTP.DELETE
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for delete item.\n", err)
		ctx.sendError(400, "Inavlid url for deleting item.")
	}
	sql.DeleteItem(uint32(id))
}

func GetItemById(ctx *Context) { //HTTP.GET
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for getting item.\n", err)
		ctx.sendError(400, "Inavlid url for getting item.")
	}
	ctx.sendAnswer(sql.GetItemById(uint32(id)))
}

func GetItemAll(ctx *Context) { //HTTP.GET
	ctx.sendAnswer(sql.GetItemAll())
}
