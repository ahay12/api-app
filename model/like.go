package model

type Like struct {
	UserID    uint `json:"user_id" gorm:"primary_key"`
	ArticleID uint `json:"article_id"`
}
