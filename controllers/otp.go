package controllers

import (
	"fmt"
	"os"

	"github.com/MohamedmuhsinJ/shopify/database"
	"github.com/MohamedmuhsinJ/shopify/models"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	Open "github.com/twilio/twilio-go/rest/verify/v2"
)

var Client *twilio.RestClient

func SentOtp(c *gin.Context) {
	AID := os.Getenv("Account_sid")
	AToken := os.Getenv("Auth_Token")
	SSID := os.Getenv("SSID")
	Mob := c.Query("number")
	res := CheckNum(Mob)
	if !res {
		c.JSON(400, gin.H{
			"errr": "Number doesnot exists",
		})
		return
	}
	Client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   AID,
		Password:   AToken,
		AccountSid: AID,
	})
	params := &Open.CreateVerificationParams{}
	params.SetTo("+91" + Mob)
	params.SetChannel("sms")

	resp, err := Client.VerifyV2.CreateVerification(SSID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		c.JSON(200, gin.H{
			"status": "true",
			"msg":    "OTP sent successfully..!!",
		})
	}
}
func CheckNum(num string) bool {
	var user models.User
	database.Db.First(&user, "phone=?", num)
	return num == user.Phone

}

func CheckOtp(c *gin.Context) {
	var user models.User
	phone := c.Query("phone")
	code := c.Query("code")
	database.Db.First(&user, "phone=?", phone)
	AID := os.Getenv("Account_sid")
	AToken := os.Getenv("Auth_Token")
	SSID := os.Getenv("SSID")

	Client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: AID,
		Password: AToken,
		//AccountSid: AID,
	})
	params := &Open.CreateVerificationCheckParams{}
	params.SetTo("+91" + user.Phone)
	params.SetCode(code)
	res, err := Client.VerifyV2.CreateVerificationCheck(SSID, params)
	if err != nil {
		c.JSON(400, gin.H{
			"err":  err.Error(),
			"code": code,
			// "mod":  user,
		})
		c.Abort()
		return
	} else if *res.Status == "approved" {
		// tokenString, err := GenerateToken(user.Email)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "failed to create tokenString",
		// 	})

		// 	return
		// }

		// token := tokenString["Token"]
		// c.SetSameSite(http.SameSiteLaxMode)
		// c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
		// c.JSON(202, gin.H{
		// 	"message": "verified",
		// 	"token":   tokenString,
		// })

	} else {
		c.JSON(400, gin.H{
			"Error": "otp is invalid",
		})
	}
}
