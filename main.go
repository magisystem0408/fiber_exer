package main

import (
	"fiber_first/database"
	"fiber_first/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	app :=fiber.New()


	//ブラウザブロックで使用する
	app.Use(cors.New(cors.Config{
		//全てのフロントエンドのエンドポイントからアクセスを可能にする
			AllowCredentials: true,
		}))
	routes.Setup(app)

	app.Listen(":3000")



}