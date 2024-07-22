package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model  `json:"gorm.Model"`
	Image       string  `gorm:"column:image;size:100;null" form:"image" json:"image" json:"image,omitempty"`
	NameProduct string  `gorm:"column:name_product;size:100;not null" form:"name_product" json:"name_product" json:"nameProduct,omitempty"`
	Price       float64 `gorm:"column:price;size;100;not null" form:"price" json:"price,omitempty"`
	Stock       uint    `gorm:"column:stock;size;100;not null" form:"stock" json:"stock,omitempty"`
	Discount    float64 `gorm:"column:discount;size;100;not null" form:"discount" json:"discount,omitempty"`
}
