package main

import (
	"log"
	"net/http"

	"github.com/adarsh2858/fiber-crm/pkg/handlers"
	"github.com/adarsh2858/fiber-crm/pkg/models"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

const (
	collectionName = "leads"
)

func setupRoutes(app *fiber.App) {
	app.Get("/crm/users", handlers.GetUsers)
	app.Get("/crm/user/:id", handlers.GetUser)
	app.Post("/crm/user", handlers.AddUser)
	app.Put("/crm/user/:id", handlers.UpdateUser)
	app.Delete("/crm/user/:id", handlers.RemoveUser)

	app.Get("/leads", func(c *fiber.Ctx) {
		query := bson.D{{}}

		cursor, err := models.Mg.Db.Collection(collectionName).Find(c.Context(), query)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}

		var leads []models.Lead = make([]models.Lead, 0)

		if err := cursor.All(c.Context(), &leads); err != nil {
			c.Status(400).SendString(err.Error())
		}
		c.JSON(leads)
	})
	app.Post("/lead", func(c *fiber.Ctx) {
		collection := models.Mg.Db.Collection(collectionName)
		parsedLeadData := new(models.Lead)

		if err := c.BodyParser(parsedLeadData); err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		parsedLeadData.ID = ""

		insertedLead, err := collection.InsertOne(c.Context(), parsedLeadData)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
			return
		}

		filter := bson.D{{Key: "_id", Value: insertedLead.InsertedID}}
		createdLead := collection.FindOne(c.Context(), filter)

		l := &models.Lead{}
		createdLead.Decode(l)
		c.Status(http.StatusCreated).JSON(l)
	})
	app.Put("/lead/:id", func (c *fiber.Ctx) {
		idParam := c.Params("id")

		leadID, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		lead := new(models.Lead)
		err = c.BodyParser(lead)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		query := bson.D{{Key:"_id", Value:leadID}}
		update := bson.D{
			{Key: "$set",
			Value: bson.D{
				{Key: "name", Value:lead.Name},
				{Key: "capital", Value:lead.Capital},
				{Key: "age", Value:lead.Age},
			},
		},}

		err = models.Mg.Db.Collection(collectionName).FindOneAndUpdate(c.Context(), query, update).Err()
		if err != nil {
			if err == models.ErrNoDocuments {
				c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			c.Status(500).SendString(err.Error())
		}

		c.Status(http.StatusOK).JSON(lead)
	},)
	app.Delete("/lead/:id", func(c *fiber.Ctx){
		idParam := c.Params("id")
		leadID, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		query := bson.D{{Key:"_id", Value: leadID}}
		result, err := models.Mg.Db.Collection(collectionName).DeleteOne(c.Context(), &query)
		if err != nil {
			c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		if result.DeletedCount < 1 {
			c.SendStatus(http.StatusNotFound)
		}

		c.Status(http.StatusOK).JSON("record deleted")
	})

	// add twilio otp verification functions and redirect to the service layer for adding logic
	app.Post("/send-otp", handlers.SendOtp)
	app.Post("/verify-otp", handlers.VerifyOtp)
}

func main() {
	app := fiber.New()

	models.InitDatabase()
	setupRoutes(app)

	if err := models.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(3000))
	defer models.DbConn.Close()
}
