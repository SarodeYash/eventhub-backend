package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "main.go/Util"
)

func Authentication(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization header missing"})
		return
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Authorization format"})
		return
	}

	token := authHeader[7:]

	Uid, err := util.ValidateJWTToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Cannot authorise", "error": err.Error()})
		return
	}

	context.Set("UserId", Uid)
	context.Next()
}
