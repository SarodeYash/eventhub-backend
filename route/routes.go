package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/middleware"
)

func RegisterRoute(server *gin.Engine) {
	//Aceepts Request from client at /event and handle it with Event_Handler Method
	server.GET("/event", Event_handler)
	//Dynamic Path handler with ":" we can get any value of Id dynamically
	server.GET("/event/:Id", Event_handlerByUserId)
	//Delete's Event from the databse with the mentioned ID
	//Accept new users and store there Credenttials in database
	server.POST("/event/signup", Event_HandlerNewUser)
	//Aceept User Login at this request
	server.POST("/login", Event_HandlerLogin)

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authentication)
	authenticate.DELETE("/event/delete/:Id", Event_HandlerDeleteEvent)
	//Accepts requeted data from Client side at /event and handles or Operate it wiht CreateEvent Method
	authenticate.POST("/event", CreateEvent)
	//Updates data mentioned at the specified user Id
	authenticate.PUT("/event/update/:Id", Event_HandlerUpdtaeEvent)
	//User Event Registration
	authenticate.POST("/event/register/:Id", EventRegistartion)
	//User Event Registration Canselation
	authenticate.DELETE("/event/cancel/:Id", EventRegistartionCans)

}
