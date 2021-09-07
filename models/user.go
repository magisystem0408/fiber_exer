package models


//後ろに付け加えることで表示形式を変更することができる

type User struct {
	Id       uint `json:"id"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
}

// ハイフンで中身を見えないようにすることができる

