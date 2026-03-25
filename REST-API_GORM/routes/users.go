package routes

import (
	"learning/rest-api_gorm/models"
	"learning/rest-api_gorm/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "user created successfully")
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, "not parse data")
		return
	}

	err = user.ValidateUser()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.Generatetoken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login sucessfully", "token": token})

}
