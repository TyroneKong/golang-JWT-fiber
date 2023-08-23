package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ID           uint    `json:"id" gorm:"primaryKey"`
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
