package main

import (
	"fmt"
	"log"
	"net/http"
	"webServer/config"
	"webServer/handlers"
	"webServer/repository"
	"webServer/service"
)

func main() {
	db := config.ConnectDB()
	repo := &repository.AccountRepository{db}
	service := &service.AccountService{repo}
	handlers := handlers.AccountHandlers{service}

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/create", handlers.CreateHandler)
	http.HandleFunc("/deposit", handlers.DepositHandler)
	http.HandleFunc("/withdraw", handlers.WithdrawHandler)
	checkBalanceHandler(handlers)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkBalanceHandler(handlers handlers.AccountHandlers) {
	http.HandleFunc("/check-balance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.BalanceHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CheckBalanceHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
