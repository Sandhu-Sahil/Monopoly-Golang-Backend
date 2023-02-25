package routes

import (
	"monopoly-Sandhu-Sahil/controllers"

	"github.com/gin-gonic/gin"
)

type RouterServiceProperty struct {
	Propertycontroller controllers.PropertyController
}

func NewRouterServiceProperty(propertyController controllers.PropertyController) RouterServiceProperty {
	return RouterServiceProperty{
		Propertycontroller: propertyController,
	}
}

func (rsp *RouterServiceProperty) RegisterPropertyRoutes(rg *gin.RouterGroup) {
	propertyRoute := rg.Group("/property")

	propertyRoute.POST("/add", rsp.Propertycontroller.CreateProperty)
	propertyRoute.GET("/getAll", rsp.Propertycontroller.GetProperties)
}
