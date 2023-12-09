package models

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"Phone"`
}

func FetchUsers(c *fiber.Ctx) ([]User, error) {
	var users []User
	DbConn.Find(&users)
	c.JSON(users)

	return []User{}, nil
}

func FetchUser() (*User, error) {
	// id := c.Params("id")
	return &User{}, nil
}

func AddUser() (*User, error) { return &User{}, nil }

func UpdateUser() (*User, error) { return &User{}, nil }

func DeleteUser() error { return nil }
