package models

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func FetchUsers(c *fiber.Ctx) ([]User, error) {
	var users []User
	DbConn.Find(&users)
	c.JSON(users)

	return []User{}, nil
}

func FetchUser(c *fiber.Ctx) (*User, error) {
	id := c.Params("id")
	var u User
	DbConn.Find(&u, id)
	c.JSON(u)

	return &User{}, nil
}

func AddUser(c *fiber.Ctx) (*User, error) {
	fmt.Println("JE")
	var u User
	if err := c.BodyParser(&u); err != nil {
		log.Print(err)
		return nil, nil
	}
	fmt.Println("JE")
	fmt.Println(u)

	DbConn.Create(&u)
	c.JSON(u)

	return &User{}, nil
}

func UpdateUser(c *fiber.Ctx) (*User, error) {
	id := c.Params("id")
	var u User
	if err := c.BodyParser(&u); err != nil {
		log.Print(u)
		return nil, nil
	}

	var existingUser User

	DbConn.First(&existingUser, id)
	DbConn.Model(&existingUser).Update(u)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return nil, nil
	// }

	c.JSON(u)

	return &User{}, nil
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var u User
	DbConn.First(&u, id)
	DbConn.Delete(u)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return nil
	// }
	c.JSON("deleted user")

	return nil
}
