package controllers

import (
	"e-commerce-golang/database"
	"e-commerce-golang/models"

)

func GetComments() []models.Comment {
	var comments []models.Comment
	res := database.GetDB().Find(&comments)
	if res.Error != nil {
		panic(res.Error)
	}
	// fmt.Println(comments)
	return comments
}

func GetCommentsById(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	//res := database.GetDB().(&comments, models.Comment{IdItem: id})
	res := database.GetDB().Where(&models.Comment{
		ItemRefer: id,
	}).Find(&comments)
	return comments, res.Error
}
