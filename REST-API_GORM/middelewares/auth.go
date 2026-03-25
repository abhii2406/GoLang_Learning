package middelewares

import (
	"fmt"
	"learning/rest-api_gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")
	fmt.Println("token:", token)
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
		return
	}
	uid, err := utils.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.Set("userId", uid)
	ctx.Next()
}
