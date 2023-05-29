package handlers

import (
	"e-commerce-golang/controllers"
	"e-commerce-golang/database"
	"e-commerce-golang/middle"
	"e-commerce-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItems(cnt *gin.Context) {
	items := controllers.GetItems()
	//items := controllers.Get()
	cnt.IndentedJSON(http.StatusOK, items)
}

// with comments
func GetItemsWithComments(cnt *gin.Context) {
	items := controllers.GetWithComments()
	//items := controllers.Get()
	cnt.IndentedJSON(http.StatusOK, items)
}

func GetItem(cnt *gin.Context) {
	id := cnt.Param("id")
	item, err := controllers.GetItemById(middle.ConvertToUint(id))
	if err != nil {
		cnt.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}
	cnt.IndentedJSON(http.StatusOK, item)
}

func Filter(cnt *gin.Context) {
	based := cnt.Query("based")
	order := cnt.Query("order")
	if based == "price" {
		items := controllers.GetSortedItemsByPrice(order)
		cnt.IndentedJSON(http.StatusOK, items)
	} else if based == "rating" {
		items := controllers.GetSortedItemsByRating(order)
		cnt.IndentedJSON(http.StatusOK, items)
	} else {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
	}
}

func AddItem(cnt *gin.Context) {
	var newItem models.Item
	if err := cnt.BindJSON(&newItem); err != nil {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "adadsads"})
		panic(err)
	}
	database.DB.Create(&newItem)

	cnt.IndentedJSON(http.StatusCreated, newItem)
}
func DeleteItem(cnt *gin.Context) {
	id := cnt.Param("id")
	item, err := controllers.GetItemById(middle.ConvertToUint(id))
	if err != nil {
		cnt.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found!"})
		return
	}
	database.DB.Delete(item)
}
