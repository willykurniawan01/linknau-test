package router

import (
	"github.com/gin-gonic/gin"
	"github.com/willykurniawan01/linknau-test/app/controllers"
	"github.com/willykurniawan01/linknau-test/app/middleware"
)

func Route(router *gin.Engine) {
	// Middleware
	jwtMiddleware := new(middleware.JwtMiddleware)

	// Controller
	authController := new(controllers.AuthController)
	userController := new(controllers.UserController)

	router.POST("/login", authController.Login)
	router.GET("/user/profile", jwtMiddleware.VerifyJwt, userController.GetProfile)
}
