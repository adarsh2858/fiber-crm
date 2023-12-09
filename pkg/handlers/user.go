package handlers

import (
	"encoding/json"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"Phone"`
}

func FetchUsers() ([]User, error) { return []User{}, nil }

func FetchUser() (*User, error) { return &User{}, nil }

func AddUser() (*User, error) { return &User{}, nil }

func UpdateUser() (*User, error) { return &User{}, nil }

func DeleteUser() error { return nil }
