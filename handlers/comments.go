package handlers

import (
	"e-commerce-golang/controllers"
	"e-commerce-golang/database"
	"e-commerce-golang/middle"
	"e-commerce-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCommentaries(cnt *gin.Context) {
	itemId := cnt.Param("id")
	comments, err := controllers.GetCommentsById(middle.ConvertToUint(itemId))
	if err != nil {
		cnt.IndentedJSON(http.StatusNotFound, gin.H{"message": "comments not found"})
		return
	}
	cnt.IndentedJSON(http.StatusOK, comments)
}
func GetComments(cnt *gin.Context) {
	comments := controllers.GetComments()
	cnt.IndentedJSON(http.StatusOK, comments)
}

func SetComment(cnt *gin.Context) {
	itemId := cnt.Param("id")
	var newComment models.Comment
	if err := cnt.BindJSON(&newComment); err != nil {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
		panic(err)
	}
	newComment.ItemRefer = middle.ConvertToUint(itemId)
	database.DB.Create(&newComment)
	controllers.SetRating(newComment.ItemRefer)
	cnt.IndentedJSON(http.StatusCreated, newComment)
}
