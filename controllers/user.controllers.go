package controllers

import (
	"monopoly-Sandhu-Sahil/models"
	"monopoly-Sandhu-Sahil/services"
	"monopoly-Sandhu-Sahil/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := uc.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login verified", "token": token})
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	valid := services.IsPasswordValid(user.Password)
	if valid != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password", "message": "Password must contain UPPER CASE, LOWER CASE, SPECIAL CHARACTER, NUMBER and LENGTH>7"})
		return
	}

	token, err := uc.UserService.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration success", "token": token})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "0" {
		id = user_id
	}

	u, err := uc.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
