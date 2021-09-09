package controllers

import (
	"fiber_first/database"
	"fiber_first/models"
	"fiber_first/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//読み込みとエラーハンドリング
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//確認用パスワードと一致しなかった時のエラーハンドリング
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	//バイト変換してあげる必要がある
	//password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		First:    data["first_name"],
		Last:     data["last_name"],
		Email:    data["email"],
	}

	user.SetPassword(data["password"])

	//データベースへのコミット
	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User

	//メールアドレスクエリで検索中
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	//パスワード検証
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//GenerateJWTでjwtトークンにエンコードした時に、シークレットキーが発行される
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))


	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//クッキーをセットする時に使用する
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}


func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	//tokenの内容をパースしている
	id, _ := util.ParseJwt(cookie)
	//claimsのissuerのところにidが格納されている
	//データベースのIDを参照しに行く
	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}


func Logout(c *fiber.Ctx) error {

	//クッキーは削除するのではなく空の値を生成する
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	//クッキーをセットする時に使用する
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
