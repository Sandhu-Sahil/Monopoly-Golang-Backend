package main

import (
	"context"
	"fmt"
	"log"
	"monopoly-Sandhu-Sahil/controllers"
	"monopoly-Sandhu-Sahil/routes"
	"monopoly-Sandhu-Sahil/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	rs          routes.RouterService
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
	// validate    *validator.Validate
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(os.Getenv("DATABASE"))
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database(os.Getenv("DATABASE_NAME")).Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	rs = routes.NewRouterService(uc)
	server = gin.Default()

	// validate = validator.New()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/api-v1")
	rs.RegisterLoginRoutes(basepath)

	jwtpath := server.Group("/api-v2")
	rs.RegisterJwtCheckRoutes(jwtpath)

	log.Fatal(server.Run(":8080"))
}
