package models

import "golang.org/x/crypto/bcrypt"

//後ろに付け加えることで表示形式を変更することができる

type User struct {
	Id       uint   `json:"id"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Email    string `gorm:"unique"`
// ハイフンで中身を見えないようにすることができる
	Password []byte `json:"-"`
}

//パスワード生成用の関数
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password =hashedPassword
}

func (user *User) ComparePassword(password string)error {
	return bcrypt.CompareHashAndPassword(user.Password,[]byte(password))

}

