package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/willykurniawan01/linknau-test/app/helpers"
)

type UserController struct{}

func (ac *UserController) GetProfile(c *gin.Context) {
	helpers.ResponseApi(c, "SUCCESS", gin.H{
		"name":            "Willy kurniawan",
		"age":             24,
		"profile_picture": "",
		"phone_number":    "081363810321",
	})
}
