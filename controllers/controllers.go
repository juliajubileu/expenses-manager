package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juliajubileu/expenses-manager/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Incomes)
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, income := range models.Incomes {
		if strconv.Itoa(income.Id) == id {
			json.NewEncoder(w).Encode(income)
		}
	}
}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Expenses)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, expense := range models.Expenses {
		if strconv.Itoa(expense.Id) == id {
			json.NewEncoder(w).Encode(expense)
		}
	}
}

