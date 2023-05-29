package handlers

import (
	"e-commerce-golang/controllers"
	"e-commerce-golang/database"
	"e-commerce-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registration(cnt *gin.Context) {
	var newUser models.User
	if err := cnt.BindJSON(&newUser); err != nil {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	user := controllers.SearchUser(newUser.Email)
	if user == (models.User{}) {
		database.DB.Create(&newUser)
		cnt.IndentedJSON(http.StatusCreated, newUser)
	} else {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email is already in use"})
		return
	}
}

func Authorisation(cnt *gin.Context) {
	var user models.User
	if err := cnt.BindJSON(&user); err != nil {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect login or password"})
		return
	}
	findUser := controllers.SearchUser(user.Email)
	if findUser == (models.User{}) {
		cnt.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no such user"})
		return
	} else if findUser.Password == user.Password {
		cnt.IndentedJSON(http.StatusOK, gin.H{"message": "Authorization was successful"})
		return
	} else {
		cnt.IndentedJSON(http.StatusForbidden, gin.H{"message": "Incorrect login or password"})
	}
}
