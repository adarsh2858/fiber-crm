package handlers

import (
	"os"
	"fmt"
	"log"
	"net/http"
	// "encoding/json"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	twilioVerify "github.com/twilio/twilio-go/rest/verify/v2"
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
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUM"))
	params.SetBody("Hello world")
	fmt.Println(params)

	// response, err := twilioClient.Api.CreateMessage(params)
	// if err != nil {
	// 	c.Status(http.StatusBadRequest).SendString(err.Error())
	// 	return
	// }
	// res, _ := json.Marshal(*response)
	// fmt.Printf("%v", string(res))

	verificationDetails  := &twilioVerify.CreateVerificationParams{}
	verificationDetails.SetTo(otpRequest.PhoneNumber)
	verificationDetails.SetChannel("sms")

	resp, err := twilioClient.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICES_ID"), verificationDetails)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}
	log.Print(*resp.Sid)

	c.Status(http.StatusOK).JSON("sent otp")
}

func VerifyOtp(c *fiber.Ctx){
	var v VerifyOtpRequest
	if err := c.BodyParser(&v); err != nil {
		c.Status(400).SendString(err.Error())
		return
	}

	params := &twilioVerify.CreateVerificationCheckParams{}
	params.SetTo(v.PhoneNumber)
	params.SetCode(v.Code)

	resp, err := twilioClient.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICES_ID"),params)
	if err != nil {
		c.Status(400).SendString(err.Error())
		return
	}

	if *resp.Status != "approved" {
		c.Status(404).JSON("invalid otp")
		return
	}

	c.Status(200).JSON("verified otp")
}
