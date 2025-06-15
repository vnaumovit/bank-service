package service

import (
	"fmt"
	"webServer/repository"
)

type AccountService struct {
	Repo *repository.AccountRepository
}

func (s *AccountService) Create(name string) error {
	return s.Repo.CreateAccount(name)
}

func (s *AccountService) Deposit(name string, amount float64) error {
	return s.Repo.UpdateBalance(name, amount)
}

func (s *AccountService) Withdraw(name string, amount float64) error {
	balance, err := s.Repo.GetBalance(name)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("insufficient funds, name='%s' amount='%.2f' balance'%.2f'", name, amount, balance)
	}
	return s.Repo.UpdateBalance(name, -amount)
}

func (s *AccountService) Balance(name string) (float64, error) {
	return s.Repo.GetBalance(name)
}
