package controllers

import (
	"fiber_first/database"
	"fiber_first/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//読み込みとエラーハンドリング
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//確認用パスワードと一致しなかった時のエラーハンドリング
	if data["password"]  !=data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"password do not match",
		})
	}

	//バイト変換してあげる必要がある
	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]),14)



	user := models.User{
		First: data["first_name"],
		Last: data["last_name"],
		Email: data["email"],
		Password: password,
	}

	//データベースへのコミット
	database.DB.Create(&user)

	return c.JSON(user)
}




func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err :=c.BodyParser(&data); err!=nil{
		return err
	}
	var user models.User


	//メールアドレスクエリで検索中
	database.DB.Where("email = ?",data["email"]).First(&user)

	if user.Id ==0{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"not found",
		})
	}

	//パスワード検証
	if err :=bcrypt.CompareHashAndPassword(user.Password,[]byte(data["password"]));err !=nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"incorrect password",
		})
	}

	return c.JSON(user)
}
