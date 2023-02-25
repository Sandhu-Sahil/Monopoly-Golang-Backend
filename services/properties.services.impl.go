package services

import (
	"context"
	"monopoly-Sandhu-Sahil/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyServiceImpl struct {
	propertycollection *mongo.Collection
	ctx                context.Context
}

func NewPropertyService(propertyCollection *mongo.Collection, ctx context.Context) PropertyService {
	return &PropertyServiceImpl{
		propertycollection: propertyCollection,
		ctx:                ctx,
	}
}

func (p *PropertyServiceImpl) CreteNewProperty(property *models.Property) (string, error) {
	propertyCreated, err := p.propertycollection.InsertOne(p.ctx, property)

	if err != nil {
		return "", err
	}
	id := propertyCreated.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (p *PropertyServiceImpl) GetAllProperties() ([]models.Property, error) {
	var properties []models.Property

	cursor, err := p.propertycollection.Find(p.ctx, bson.M{})
	if err != nil {
		return properties, err
	}

	defer cursor.Close(p.ctx)
	for cursor.Next(p.ctx) {
		var property models.Property
		cursor.Decode(&property)
		properties = append(properties, property)
	}
	if err = cursor.Err(); err != nil {
		return properties, err
	}

	return properties, err
}
