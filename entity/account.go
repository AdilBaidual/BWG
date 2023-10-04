package entity

type Account struct {
	Id            int     `json:"account" db:"id"`
	CurrencyCode  string  `json:"currency_code" db:"currency_code" binding:"required"`
	ActiveBalance float64 `json:"active_balance" db:"active_balance"`
	FrozenBalance float64 `json:"frozen_balance" db:"frozen_balance"`
	UserId        *int    `json:"user_id" db:"user_id" binding:"required"`
}
