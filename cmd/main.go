package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adarsh2858/fiber-crm/pkg/handlers"
	"github.com/adarsh2858/fiber-crm/pkg/models"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dBName         = "fiber-hrms"
	collectionName = "leads"
	mongoURI       = "mongodb://localhost:27017/" + dBName
)

var mg *MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type Lead struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Capital float64 `json:"capital"`
	Age     uint8   `json:"age"`
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

	mg = &MongoInstance{
		Client: client,
		Db:     dB,
	}
	return err
}

func setupRoutes(app *fiber.App) {
	app.Get("/crm/users", handlers.GetUsers)
	app.Get("/crm/user/:id", handlers.GetUser)
	app.Post("/crm/user", handlers.AddUser)
	app.Put("/crm/user/:id", handlers.UpdateUser)
	app.Delete("/crm/user/:id", handlers.RemoveUser)
}

func main() {
	app := fiber.New()

	models.InitDatabase()
	setupRoutes(app)

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app.Get("/leads", func(c *fiber.Ctx) {
		query := bson.D{{}}

		cursor, err := mg.Db.Collection(collectionName).Find(c.Context(), query)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}

		var leads []Lead = make([]Lead, 0)

		if err := cursor.All(c.Context(), leads); err != nil {
			c.Status(400).SendString(err.Error())
		}
		c.JSON(leads)
	})

	app.Post("/lead", func(c *fiber.Ctx) {
		collection := mg.Db.Collection(collectionName)
		parsedLeadData := new(Lead)

		if err := c.BodyParser(parsedLeadData); err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		parsedLeadData.ID = ""
		log.Print(parsedLeadData)
		log.Print(parsedLeadData)

		insertedLead, err := collection.InsertOne(c.Context(), parsedLeadData)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
			return
		}
		log.Print(insertedLead)

		filter := bson.D{{Key: "_id", Value: insertedLead.InsertedID}}
		log.Print(filter)
		createdLead := collection.FindOne(c.Context(), filter)
		log.Print(createdLead)

		l := &Lead{}
		createdLead.Decode(l)
		c.Status(http.StatusCreated).JSON(l)
	})

	log.Fatal(app.Listen(3000))
	defer models.DbConn.Close()
}
