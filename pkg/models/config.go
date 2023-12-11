package models

import (
	"fmt"
	"time"
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

const (
	mongoURI       = "mongodb://localhost:27017/" + dBName
	dBName         = "fiber-leads"
)

var ErrNoDocuments = mongo.ErrNoDocuments


var Mg *MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	dB := client.Database(dBName)

	Mg = &MongoInstance{
		Client: client,
		Db:     dB,
	}
	return err
}

