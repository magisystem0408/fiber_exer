package models

type User struct {
	Id       uint
	First    string
	Last     string
	Email    string `gorm:"unique"`
	Password []byte
}
