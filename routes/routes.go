package routes

import (
	"fmt"

	"github.com/colombohenrique/rollin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	porta := ":8080" //to specify the API port

	r := gin.Default()

	//Home
	r.GET("/", controllers.Home) //Page containing API documentation

	//Events
	r.GET("/api/event", controllers.GetAllEvents)
	r.GET("/api/event/:id", controllers.GetEvent)
	r.GET("/api/event/name/:eventName", controllers.SearchEventByName)
	r.POST("/api/event", controllers.AddNewEvent)
	r.DELETE("/api/event/:id", controllers.DeleteEvent)
	r.PATCH("/api/event/:id", controllers.EditEvent)

	//Items
	r.GET("/api/item", controllers.GetAllItems)
	r.GET("/api/item/:id", controllers.GetItem)
	r.POST("/api/item", controllers.AddNewItem)
	r.DELETE("/api/item/:id", controllers.DeleteItem)
	r.PATCH("/api/item/:id", controllers.EditItem)

	//Users
	r.GET("/api/user", controllers.GetAllUsers)
	r.GET("/api/user/:id", controllers.GetUser)
	r.POST("/api/user", controllers.AddNewUser)
	r.PATCH("/api/user/:id", controllers.EditUser)

	//r.Run()
	r.Run(porta)
	fmt.Println("API rodando na porta: " + porta)
}
