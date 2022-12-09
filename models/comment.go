package models

type Comment struct {
	Id uint `json:"id"`
	Comment string `json:"comment"`
	ForumId uint `json:"forum_id"`
	UserId uint `json:"user_id"`
	Username string `json:"username"`
}