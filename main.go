package main

import (
	"fmt"

	"github.com/juliajubileu/expenses-manager/database"
	"github.com/juliajubileu/expenses-manager/routes"
)

func main() {
	database.ConnectToDatabase()
	fmt.Println("Initializing Server")
	routes.HandleRequest()
}
