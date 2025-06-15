package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"webServer/service"
)

type AccountController struct {
	Service *service.AccountService
}

func (c *AccountController) Create(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Empty name in request", http.StatusBadRequest)
	}
	err := c.Service.Create(name)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Account created: %s\n", name)
}

func (c *AccountController) Deposit(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	amountStr := r.URL.Query().Get("amount")
	amount, _ := strconv.ParseFloat(amountStr, 64)
	err := c.Service.Deposit(name, amount)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Deposited %.2f to %s\n", amount, name)
}
