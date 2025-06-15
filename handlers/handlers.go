package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"webServer/service"
)

type AccountHandlers struct {
	Service *service.AccountService
}

var templates = template.Must(template.ParseFiles(
	"templates/home.html",
	"templates/create.html",
	"templates/deposit.html",
	"templates/withdraw.html",
	"templates/balance.html",
	"templates/check-balance.html",
))

func (h *AccountHandlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func (h *AccountHandlers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.ExecuteTemplate(w, "create.html", nil)
	case http.MethodPost:
		name := r.FormValue("name")
		if name == "" {
			templates.ExecuteTemplate(w, "create.html", MessageData{Error: "Account name is required"})
			return
		}
		err := h.Service.Create(name)
		if err != nil {
			templates.ExecuteTemplate(w, "create.html", MessageData{Error: err.Error()})
			return
		}
		templates.ExecuteTemplate(w, "create.html", MessageData{Message: "Account created successfully!"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type MessageData struct {
	Message string
	Error   string
}

func (h *AccountHandlers) DepositHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.ExecuteTemplate(w, "deposit.html", nil)
	case http.MethodPost:
		name := r.FormValue("name")
		amountStr := r.FormValue("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount <= 0 {
			templates.ExecuteTemplate(w, "deposit.html", MessageData{Error: "Invalid amount"})
			return
		}

		err = h.Service.Deposit(name, amount)
		if err != nil {
			templates.ExecuteTemplate(w, "deposit.html", MessageData{Error: err.Error()})
			return
		}
		templates.ExecuteTemplate(w, "deposit.html", MessageData{Message: "Deposit successful!"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AccountHandlers) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.ExecuteTemplate(w, "withdraw.html", nil)
	case http.MethodPost:
		name := r.FormValue("name")
		amountStr := r.FormValue("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount <= 0 {
			templates.ExecuteTemplate(w, "withdraw.html", MessageData{Error: "Invalid amount"})
			return
		}
		err = h.Service.Withdraw(name, amount)
		if err != nil {
			templates.ExecuteTemplate(w, "withdraw.html", MessageData{Error: err.Error()})
			return
		}
		templates.ExecuteTemplate(w, "withdraw.html", MessageData{Message: "Withdrawal successful!"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type BalanceData struct {
	Name    string
	Balance float64
	Error   string
}

func (h *AccountHandlers) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "check-balance.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Template error:", err)
	}
}

func (h *AccountHandlers) CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/check-balance", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		templates.ExecuteTemplate(w, "check-balance.html", BalanceData{Error: "Please enter an account name."})
		return
	}

	balance, err := h.Service.Balance(name)
	if err != nil {
		templates.ExecuteTemplate(w, "check-balance.html", BalanceData{Error: "Account not found or error: " + err.Error()})
		return
	}

	data := BalanceData{
		Name:    name,
		Balance: balance,
	}

	// Show the balance page
	err = templates.ExecuteTemplate(w, "balance.html", data)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}
