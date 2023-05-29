package controllers

import (
	"e-commerce-golang/database"
	"e-commerce-golang/models"
	"github.com/jinzhu/gorm"
)

func GetItems() []models.Item {
	var items []models.Item
	res := database.GetDB().Find(&items)
	if res.Error != nil {
		panic(res.Error)
	}
	return items
}

func GetItemById(id uint) (models.Item, error) {
	var item models.Item
	res := database.GetDB().Find(&item, models.Item{
		Model: gorm.Model{ID: id},
	})
	return item, res.Error
}

func GetSortedItemsByRating(order string) []models.Item {
	var items []models.Item
	if order == "asc" {
		database.GetDB().Order("price asc, rating").Find(&items)
	} else if order == "desc" {
		database.GetDB().Order("price desc, rating").Find(&items)
	}
	return items
}

func GetSortedItemsByPrice(order string) []models.Item {
	var items []models.Item
	if order == "asc" {
		database.GetDB().Order("price asc, price").Find(&items)
	} else if order == "desc" {
		database.GetDB().Order("price desc, price").Find(&items)
	}
	return items
}

func GetWithComments() []models.Item {
	var items []models.Item
	database.GetDB().Preload("Comments").Find(&items)
	return items
}

func SetRating(id uint) {
	//id := newComment.ItemRefer
	var rating float64
	var cnt float64
	comments, err := GetCommentsById(id)
	if err != nil {
		panic(err)
	}
	for i := range comments {
		rating += comments[i].Rating
		cnt++
	}
	item, err := GetItemById(id)
	if err != nil {
		panic(err)
	}
	rating = rating / cnt
	database.DB.Model(&item).Update("rating", rating)

}
