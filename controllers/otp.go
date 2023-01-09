package controllers

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TAccountSid string = os.Getenv("Account_sid")
var TAuthToken string = os.Getenv("Auth_Token")
var SSID string = os.Getenv("SSID")
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TAccountSid,
	Password: TAuthToken,
})

func Otp(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(SSID, params)
	if err != nil {
		fmt.Print("sorry unable to verify number")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("sent verification '%s' \n", *resp.Sid)
	}

}

func CheckOtp(to string) {
	var code string
	fmt.Println("Please check your phone and enter the code:")
	fmt.Scanln(&code)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(SSID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
}
