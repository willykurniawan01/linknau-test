package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/willykurniawan01/linknau-test/app/helpers"
	"github.com/willykurniawan01/linknau-test/app/requests"
)

type AuthController struct{}

func (ac *AuthController) Login(c *gin.Context) {
	var loginPayload requests.LoginPayload

	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		helpers.ResponseApi(c, "INVALID_PAYLOAD", gin.H{})
		return
	}
	// Payload Validation
	validate := validator.New()
	if err := validate.Struct(loginPayload); err != nil {
		helpers.ResponseApi(c, "INVALID_PAYLOAD", gin.H{})
		return
	}

	if loginPayload.Username != "willy" && loginPayload.Password != helpers.HashSHA256("willy") {
		helpers.ResponseApi(c, "UNAUTHORIZED", gin.H{})
	}

	user := gin.H{
		"username":     "Willy kurniawan",
		"phone_number": "081363810321",
	}

	token, err := helpers.GenerateJWT(user)
	if err != nil {
		log.Println("Error : ", err)
		helpers.ResponseApi(c, "GENERAL_ERROR", gin.H{})
		return
	}
	helpers.ResponseApi(c, "SUCCESS", gin.H{
		"token": token,
	})

}
