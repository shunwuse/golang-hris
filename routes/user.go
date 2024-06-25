package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunwuse/go-hris/controllers"
	"github.com/shunwuse/go-hris/lib"
)

type UserRoute struct {
	logger         lib.Logger
	userController controllers.UserController
}

func NewUserRoute() UserRoute {
	logger := lib.GetLogger()

	userController := controllers.NewUserController()

	return UserRoute{
		logger:         logger,
		userController: userController,
	}
}

func (r UserRoute) Setup(router *gin.Engine) {
	r.logger.Info("Setting up user routes")

	router.POST("/users", r.userController.CreateUser)
}
