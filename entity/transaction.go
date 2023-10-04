package entity

import "time"

type Transaction struct {
	Id                 int       `json:"id" db:"id"`
	CurrencyCode       string    `json:"currency_code" db:"currency_code"`
	TransactionStatus  string    `json:"transaction_status" db:"transaction_status"`
	SenderAccountID    *int      `json:"sender_account_id" db:"sender_account_id"`
	RecipientAccountID *int      `json:"recipient_account_id" db:"recipient_account_id"`
	Amount             float64   `json:"amount" db:"amount"`
	TransactionDate    time.Time `json:"transaction_date" db:"transaction_date"`
}

type Invoice struct {
	CurrencyCode       string  `json:"currency_code" db:"currency_code" binding:"required"`
	TransactionStatus  string  `json:"transaction_status" db:"transaction_status"`
	RecipientAccountID *int    `json:"recipient_account_id" db:"recipient_account_id" binding:"required"`
	Amount             float64 `json:"amount" db:"amount" binding:"required"`
}

type Withdraw struct {
	CurrencyCode       string  `json:"currency_code" db:"currency_code" binding:"required"`
	TransactionStatus  string  `json:"transaction_status" db:"transaction_status"`
	SenderAccountID    *int    `json:"sender_account_id" db:"sender_account_id" binding:"required"`
	RecipientAccountID *int    `json:"recipient_account_id" db:"recipient_account_id"`
	Amount             float64 `json:"amount" db:"amount" binding:"required"`
}
