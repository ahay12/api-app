package model

import "gorm.io/gorm"

type News struct {
	gorm.Model `json:"gorm.Model"`
	IdUsers    uint   `gorm:"column:id_users;not null" form:"id_users" json:"id_users,omitempty"`
	User       Users  `gorm:"foreignKey:IdUsers;references:ID" json:"id_user"`
	Image      string `gorm:"column:image;size:100;null" form:"image" json:"image" json:"image,omitempty"`
	Title      string `gorm:"column:name_product;size:100;not null" form:"name_product" json:"name_product" json:"nameProduct,omitempty"`
	Content    string `gorm:"column:content;size;100;not null" form:"content" json:"content,omitempty"`
	Tag        string `gorm:"column:tag;size;not null"form:"tag" json:"content,omitempty"`
}
