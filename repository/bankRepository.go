package repository

import (
	"database/sql"
)

type AccountRepository struct {
	DB *sql.DB
}

func (r *AccountRepository) CreateAccount(name string) error {
	_, err := r.DB.Exec("INSERT INTO accounts (name, balance) VALUES ($1, 0)", name)
	return err
}

func (r *AccountRepository) GetBalance(name string) (float64, error) {
	var balance float64
	err := r.DB.QueryRow("SELECT balance FROM accounts WHERE name = $1", name).Scan(&balance)
	return balance, err
}

func (r *AccountRepository) UpdateBalance(name string, delta float64) error {
	_, err := r.DB.Exec("UPDATE accounts SET balance = balance + $1 WHERE name = $2", delta, name)
	return err
}
