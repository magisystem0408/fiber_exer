package controllers

import (
	"fiber_first/database"
	"fiber_first/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(&roles)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}


	database.DB.Create(&role)
	return c.JSON(role)

}

func GetRole(c *fiber.Ctx) error {
	id,_ :=strconv.Atoi(c.Params("id"))


	role :=models.Role{
		Id: uint(id),
	}
	database.DB.Find(&role)
	return c.JSON(role)

}

func UpdateRole(c *fiber.Ctx) error {
	id,_ :=strconv.Atoi(c.Params("id"))
	role :=models.Role{
		Id: uint(id),
	}

	if err :=c.BodyParser(&role);err!=nil{
		return err
	}

	database.DB.Find(&role).Updates(role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx)error  {
	id,_ :=strconv.Atoi(c.Params("id"))
	role :=models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	return nil
}