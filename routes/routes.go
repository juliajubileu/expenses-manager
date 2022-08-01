package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliajubileu/expenses-manager/controllers"
)

func HandleRequest()  {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/incomes", controllers.GetAllIncomes).Methods("Get")
	r.HandleFunc("/api/incomes/{id}", controllers.GetIncome).Methods("Get")
	r.HandleFunc("/api/expenses", controllers.GetAllExpenses).Methods("Get")
	r.HandleFunc("/api/expenses/{id}", controllers.GetExpense).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
