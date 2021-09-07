package controllers

import (
	"fiber_first/database"
	"fiber_first/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		First:    data["first_name"],
		Last:     data["last_name"],
		Email:    data["email"],
		Password: password,
	}

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
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		//strcov：uuidからintに変換してくれる
		Issuer: strconv.Itoa(int(user.Id)),
		//jwtの有効期限
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1day
	})

	//↑でjwtトークンにエンコードした時に、シークレットキーが発行される
	token, err := claims.SignedString([]byte("secret"))
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

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	//tokenの内容をパースしている
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	//認証エラーハンドリング
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*Claims)


	//claimsのissuerのところにidが格納されている
	//データベースのIDを参照しに行く
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

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
