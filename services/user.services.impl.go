package services

import (
	"context"
	"errors"
	"monopoly-Sandhu-Sahil/models"
	"monopoly-Sandhu-Sahil/token"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	objectid, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err := u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return &models.User{}, err
	}

	userFound.Password = "**PROTECTED**"
	return userFound, nil
}

func (u *UserServiceImpl) RegisterUser(user *models.User) (string, error) {
	query := bson.D{bson.E{Key: "user_name", Value: user.UserName}}
	res, err := u.usercollection.Find(u.ctx, query)
	if err != nil {
		return "", err
	}
	// fmt.Print(res.RemainingBatchLength())
	if res.RemainingBatchLength() != 0 {
		return "", errors.New("User already existed!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.UserName = strings.TrimSpace(user.UserName)

	userCreated, err := u.usercollection.InsertOne(u.ctx, user)
	if err != nil {
		return "", err
	}
	id := userCreated.InsertedID.(primitive.ObjectID).Hex()

	token, err := token.GenerateToken(id, user.UserName, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserServiceImpl) LoginUser(user *models.User) (string, error) {
	var userFound *models.User
	query := bson.D{bson.E{Key: "user_name", Value: user.UserName}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(userFound.ID.Hex(), user.UserName, user.Email)
	if err != nil {
		return "", err
	}

	return token, err
}
