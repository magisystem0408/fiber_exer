package database

import (
	"fiber_first/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//データベースをエクスポートする
var DB *gorm.DB

func Connect() {
	database,err :=gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/gp_admin"),&gorm.Config{})

	if err !=nil{
		panic("データベースが接続できませんでした")
	}else {
		fmt.Println("接続完了")
	}

	//データベースの指定
	DB =database

	database.AutoMigrate(&models.User{})

}