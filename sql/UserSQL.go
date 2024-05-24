package sql

import (
	"golang_pr/entity"
)

func AddUser(user entity.User) {
	database.Table("users").Select("fio", "email", "password", "login", "status").Create(user)
}

func ChangeUser(id uint32, updts entity.User) {
	database.Table("users").Where("user_id = ?", id).Updates(updts)
}
func DeleteUser(id uint32) {
	var user entity.User
	database.Table("users").Delete(&user, id)
}

func GetAllUsers() []*entity.User {
	var users []*entity.User
	database.Table("users").Find(&users)
	return users
}

func GetUserById(id uint32) entity.User {
	var user entity.User
	database.Table("users").Where("user_id = ?", id).First(&user)
	return user
}
