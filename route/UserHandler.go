package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "main.go/Models"
	util "main.go/Util"
)

// Handels Post method for Registration of new user
func Event_HandlerNewUser(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Email and password"})
		return
	}
	err = user.RegisterNewUser()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Register User"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "User Registered"})
}

// Handels login Method
func Event_HandlerLogin(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid Username and Password"})
		return
	}
	err = user.Login()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Username/Password Please SignUp", "Error": err})
		return
	}
	token, err := util.GeerateJWTToken(user.Email, user.UserId) //Tocken is generated at the time of Login
	if err != nil {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{"Message": "Cannot Authorise you"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
}
