package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliajubileu/expenses-manager/database"
	"github.com/juliajubileu/expenses-manager/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	var i []models.Income
	database.DB.Find(&i)
	json.NewEncoder(w).Encode(i)
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var i models.Income
	database.DB.First(&i, id)
	json.NewEncoder(w).Encode(i)
}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	var e []models.Expense
	database.DB.Find(&e)
	json.NewEncoder(w).Encode(e)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var e models.Expense
	database.DB.First(&e, id)
	json.NewEncoder(w).Encode(e)
}

