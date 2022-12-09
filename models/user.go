package models

// struct is similar to a class
type User struct {

	Id        uint   `json:"id"`
    Name 	string `json:"name"`
    Email     string `json:"email" gorm:"unique"`
    Password  string `json:"-"` //Dont show the password in the methods
}