package controllers

import "github.com/gin-gonic/gin"

func GetAllItems(c *gin.Context) {

}

func GetItem(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}

func AddNewItem(c *gin.Context) {

}

func DeleteItem(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}

func EditItem(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}
