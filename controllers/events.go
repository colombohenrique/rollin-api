package controllers

import (
	"net/http"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	var events []models.Event
	database.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	var event models.Event
	id := c.Params.ByName("id")

	database.DB.First(&event, id)

	if event.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Event not found",
		})
		return
	}
	c.JSON(http.StatusOK, event)
}

func AddNewEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems creating new event",
			"error":  err.Error(),
		})
		return
	}
	database.DB.Create(&event)
	c.JSON(http.StatusOK, &event)
}

func DeleteEvent(c *gin.Context) {
	var event models.Event
	id := c.Params.ByName("id")

	database.DB.Delete(&event, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Event deleted",
	})
}

func EditEvent(c *gin.Context) {
	var event models.Event
	id := c.Params.ByName("id")

	database.DB.First(&event, id)

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems editing event",
			"error":  err.Error(),
		})
		return
	}

	database.DB.Model(&event).UpdateColumns(event)

	c.JSON(http.StatusOK, &event)
}

func SearchEventByName(c *gin.Context) {
	var events []models.Event
	eventName := c.Param("eventName")
	likeQuery := "%" + eventName + "%"
	database.DB.Model(&events).Where("name LIKE ?", likeQuery).Find(&events)

	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "No events found",
		})
		return
	}
	c.JSON(http.StatusOK, events)
}
