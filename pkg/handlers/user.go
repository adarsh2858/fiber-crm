package handlers

import (
	"context"
	"fmt"

	"github.com/adarsh2858/fiber-crm/pkg/models"

	"github.com/gofiber/fiber"
)

func GetUsers(c *fiber.Ctx) {
	bg := context.Background()
	fmt.Printf("%v \n", bg)
	fmt.Println("GetUsers")
	models.FetchUsers(c)
	return
}

func GetUser(c *fiber.Ctx) {
	fmt.Println("GetUser")
	return
}

func AddUser(c *fiber.Ctx) {
	fmt.Println("AddUser")
}

func UpdateUser(c *fiber.Ctx) {
	fmt.Println("UpdateUser")
}

func RemoveUser(c *fiber.Ctx) {
	fmt.Println("RemoveUser")
}
