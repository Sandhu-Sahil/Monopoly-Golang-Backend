package services

import "monopoly-Sandhu-Sahil/models"

type PropertyService interface {
	CreteNewProperty(*models.Property) (string, error)
	GetAllProperties() ([]models.Property, error)
}
