package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "main.go/Models"
)

func EventRegistartion(context *gin.Context) {
	Param_Id := context.Param("Id")
	id, err := strconv.ParseInt(Param_Id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Please Provide Valid ID"})
		return
	}
	event, err := models.GetIDEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Event ID not found"})
		return
	}

	userId := context.GetInt64("UserId")

	err = event.Registartion(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Register"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Registration Successful"})
}
func EventRegistartionCans(context *gin.Context) {
	Param_Id := context.Param("Id")
	id, err := strconv.ParseInt(Param_Id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Please Provide Valid ID"})
		return
	}
	event, err := models.GetIDEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Event ID not found"})
		return
	}
	userid := context.GetInt64("UserId")

	err = event.Cancelation(userid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Cancel"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Cancellation Successful"})

}
