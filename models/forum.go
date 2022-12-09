package models

type Forum struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ForumId"`
	UserId uint `json:"user_id"`
	Username string `json:"username"`
}