package entity

import (
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	Balance   float64   `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(id string, balance float64) *Account {
	account := &Account{
		ID:        id,
		Balance:   balance,
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) UpdateBalance(amount float64) {
	a.Balance = amount
	a.UpdatedAt = time.Now()
}
