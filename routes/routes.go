package routes

import (
	"monopoly-Sandhu-Sahil/controllers"
	"monopoly-Sandhu-Sahil/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterService struct {
	Usercontroller controllers.UserController
}

func NewRouterService(usercontroller controllers.UserController) RouterService {
	return RouterService{
		Usercontroller: usercontroller,
	}
}

func (rs *RouterService) RegisterLoginRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	// userroute.Use(middlewares.ValidationUserLogin()).POST("/login", rs.Usercontroller.Login)

	userroute.POST("/login", rs.Usercontroller.Login)
	userroute.POST("/register", rs.Usercontroller.Register)
}

func (rs *RouterService) RegisterJwtCheckRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.Use(middlewares.JwtAuthMiddleware())
	userroute.GET("/get/:id", rs.Usercontroller.GetUser)
}
