package handlers

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SendOtpRequest struct {
	PhoneNumber string `json:"phone"`
}

type VerifyOtpRequest struct {
	PhoneNumber string `json:"phone"`
	Code string `json:"code"`
}

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data string `json:"data"`
}

var twilioClient *twilio.RestClient

func init() {
	godotenv.Load()
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTHTOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	twilioClient = client
}

func SendOtp(c *fiber.Ctx){
	var otpRequest SendOtpRequest
	if err := c.BodyParser(&otpRequest); err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}
	fmt.Println(otpRequest)

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(otpRequest.PhoneNumber)
	fmt.Println(otpRequest.PhoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUM"))
	params.SetBody("Hello world")

	response, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}

	res, _ := json.Marshal(*response)
	fmt.Printf("%v", string(res))

	// twilio.SendOtp(&client.SendOtpRequest{
	// 	PhoneNumber: otpRequest.PhoneNumber,
	// })

	c.Status(http.StatusOK).JSON("sent otp")
}
func VerifyOtp(c *fiber.Ctx){
	c.JSON("verifying otp")
}
