package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "main.go/Models"
)

// Handels Get Method to Get all Events
func Event_handler(context *gin.Context) {
	GetEvent, err := models.GetAllEvents() //Used to get all the events that is stored in database
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Data Not Fetched"})
	}
	context.JSON(http.StatusOK, GetEvent)

}

// Handels Post Method to get database from client and store in the Database
func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event) //ShouldBindJSON(&event) Accepts parameters from client act as an fmt.Scan method

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Data Not fetched", "Error": err}) //If there's something error StatusBadRequest=400 will be displayed
		return
	}
	uid := context.GetInt64("UserId")
	event.UserID = uid
	err = event.Save() //data will be shared to save method in event.go

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Event Not Created"}) //Data will be saved in the data base
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created", "Event": event}) //If all ok StatusCreated=201 will be displayed along with the data
}

// Handels Get Method to get the Event Data at mentioned ID
func Event_handlerByUserId(context *gin.Context) {
	Param_Id := context.Param("Id")               //Will get us an Id Parameter from the URL in string format
	Id, err := strconv.ParseInt(Param_Id, 10, 64) //this will covert id to int64 where 10 secify we are working with decimal system and 64 specify the currenr bitsize we want
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "invalid ID Format"})
		return
	}
	GetEven_id, err := models.GetIDEvent(Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Fetch data for given ID"})
		return
	}
	context.JSON(http.StatusOK, GetEven_id)
}

// Handels Delete Method which is used to Delete Event data at mentioned ID
func Event_HandlerDeleteEvent(context *gin.Context) {
	Param_Id := context.Param("Id")
	eventID, err := strconv.ParseInt(Param_Id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid id "})
		return
	}
	event, err := models.GetIDEvent(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Fetch Data At Mentioned ID"})
		return
	}

	UserId := context.GetInt64("UserId")
	if UserId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Only Creater Can delete the Event"})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot delete Database"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Event deleted Succesfully"})

}

// Handels Update Method which is used to update the databse with the specified ID
func Event_HandlerUpdtaeEvent(context *gin.Context) {

	Event_Id := context.Param("Id")
	ActualId, err := strconv.ParseInt(Event_Id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
		return
	}

	event, err := models.GetIDEvent(ActualId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot Fetch Data At Mentioned ID"})
		return
	}
	UserId := context.GetInt64("UserId")
	if UserId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Only Creater Can update Event"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid Event Data"})
		return
	}
	updateEvent.ID = UserId
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{"Message": "Cannot Update Event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Event Updated"})

}
