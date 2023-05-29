package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Text      string  `json:"text"`
	ItemRefer uint    `json:"itemRefer"`
	Rating    float64 `json:"rating"`
}
