package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DbConn *gorm.DB

func InitDatabase() {
	// for latest gorm.io/gorm & mysql (driver reqd since server is present for it)
	// dsn := "adarsh:password@localhost:3306/fiber-crm?charset=utf8mb4&parseTime=True&loc=Local"
	// DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var err error
	DbConn, err = gorm.Open("sqlite3", "crm-fiber.db")
	if err != nil {
		fmt.Println("could not connect to database")
		panic(err)
	}
	fmt.Println("Connected to Database")
	DbConn.AutoMigrate(&User{})
	fmt.Println("Migrated")
}
