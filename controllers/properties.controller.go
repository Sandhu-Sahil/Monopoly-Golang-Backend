package controllers

import (
	"monopoly-Sandhu-Sahil/models"
	"monopoly-Sandhu-Sahil/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyController struct {
	PropertyService services.PropertyService
}

func NewProperty(propertyService services.PropertyService) PropertyController {
	return PropertyController{
		PropertyService: propertyService,
	}
}

func (pr *PropertyController) CreateProperty(ctx *gin.Context) {
	var property models.Property

	if err := ctx.ShouldBindJSON(&property); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	objectid, err := pr.PropertyService.CreteNewProperty(&property)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Property added successfully", "id": objectid})
}

func (pr *PropertyController) GetProperties(ctx *gin.Context) {
	data, err := pr.PropertyService.GetAllProperties()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "All properties fetched", "data": data})
}
