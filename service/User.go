package service

import (
	"golang_pr/entity"
	"golang_pr/sql"
	"log"
	"strconv"
)

func init() {
	CRUDS["UserCreate"] = CreateUser
	CRUDS["UserUpdate"] = UpdateUser
	CRUDS["UserDelete"] = DeleteUser
	CRUDS["Users"] = GetUserAll
	CRUDS["User"] = GetUserById
}

func CreateUser(ctx *Context) { //HTTP.POST
	var User entity.User
	User, err := GetData[entity.User](ctx)
	if err != nil {
		log.Println("Error decoding json for create user:", err)
		ctx.sendError(500, "Error decoding json for create user.")
		return
	}

	sql.AddUser(User)

	response := map[string]interface{}{"success": true}
	ctx.sendAnswer(response)
}

func UpdateUser(ctx *Context) { //HTTP.PUT
	var User entity.User
	User, err := GetData[entity.User](ctx)
	if err != nil {
		log.Println("Error of decoding json for updating user.\n", err)
		ctx.sendError(500, "Error of json decoding.")
	}
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for update user.\n", err)
		ctx.sendError(400, "Invalid url for updating user.")
	}
	sql.ChangeUser(uint32(id), User)
}

func DeleteUser(ctx *Context) { //HTTP.DELETE
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for delete user.\n", err)
		ctx.sendError(400, "Inavlid url for deleting user.")
	}
	sql.DeleteUser(uint32(id))
}

func GetUserById(ctx *Context) { //HTTP.GET
	id, err := strconv.ParseUint(ctx.Request.URL.Query()["id"][0], 10, 32)
	if err != nil {
		log.Println("Error of id reading for getting user.\n", err)
		ctx.sendError(400, "Inavlid url for getting user.")
	}
	ctx.sendAnswer(sql.GetUserById(uint32(id)))
}

func GetUserAll(ctx *Context) { //HTTP.GET
	ctx.sendAnswer(sql.GetAllUsers())
}
