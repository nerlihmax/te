package models

import "github.com/nerlihmax/te/internal/account"

type Transaction struct {
	ID          string
	Debit       int32
	Credit      int32
	Account     account.Account
	Description string
}
