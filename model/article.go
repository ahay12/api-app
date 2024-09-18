package model

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model `json:"gorm.Model"`
	Image      string `gorm:"column:image;size:100;null" form:"image" json:"image,omitempty"`
	Title      string `gorm:"column:title;size:100;not null" form:"title" json:"title,omitempty"`
	Content    string `gorm:"column:content;size:1000;not null" form:"content" json:"content,omitempty"`
	Category   string `gorm:"column:category;size:50;not null" form:"category" json:"category,omitempty"`
	Tag        string `gorm:"column:tag;size:100;null" form:"tag" json:"tag,omitempty"`
	Like       []Like `gorm:"column:like" form:"like" json:"like,omitempty"`
	View       uint   `gorm:"column:view" form:"view" json:"view,omitempty"`
}
