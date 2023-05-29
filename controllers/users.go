package controllers

import (
	"e-commerce-golang/database"
	"e-commerce-golang/models"
)

func SearchUser(target string) models.User {
	var user models.User
	database.GetDB().Where("email = ?", target).First(&user)
	return user
}
