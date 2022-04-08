package main

import (
	"fmt"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/routes"
)

func main() {
	fmt.Println("Tentando conectar...")
	database.ConnectDatabase()
	fmt.Println("Iniciando API")
	routes.HandleRequest()
	fmt.Println("API iniciada!")
}
