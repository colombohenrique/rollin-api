package controllers

import (
	"net/http"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	database.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)

}

func AddNewUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems creating new user",
			"error":  err.Error(),
		})
		return
	}

	//Data validator
	if err := models.DataValidate(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems creating new user",
			"error":  err.Error(),
		})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func EditUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	database.DB.First(&user, id)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems editing user",
			"error":  err.Error(),
		})
		return
	}

	//Data validator
	if err := models.DataValidate(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return": "Problems creating new user",
			"error":  err.Error(),
		})
		return
	}

	database.DB.Model(&user).UpdateColumns(user)

	c.JSON(http.StatusOK, &user)
}
