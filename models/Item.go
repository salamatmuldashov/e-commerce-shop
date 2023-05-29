package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name     string    `json:"name"`
	Price    string    `json:"price"`
	Comments []Comment `gorm:"foreignKey:ItemRefer"`
	Rating   float64   `json:"rating"`
}
