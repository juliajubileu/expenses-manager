package main

import (
	"fmt"

	"github.com/juliajubileu/expenses-manager/models"
	"github.com/juliajubileu/expenses-manager/routes"
)

func main() {
	models.Incomes = []models.Income {
		{Id: 1, Description: "Teste"},
	}

	fmt.Println("Initializing Server")
	routes.HandleRequest()
}
