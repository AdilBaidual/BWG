package repo

import (
	"BWG/entity"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db    *sqlx.DB
	cache *RedisRepository
}

func NewAccountPostgres(db *sqlx.DB, cache *RedisRepository) *AccountPostgres {
	return &AccountPostgres{db: db, cache: cache}
}

func (r *AccountPostgres) CreateAccount(account entity.Account) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (currency_code, user_id)"+
		" values ($1, $2) RETURNING id", accountTable)
	row := r.db.QueryRow(query, account.CurrencyCode, account.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AccountPostgres) GetAccount(id int) (entity.Account, error) {
	var account entity.Account

	cacheKey := fmt.Sprintf("account_%d", id)
	cachedData, err := r.cache.Get(cacheKey)
	if err == nil {
		if err := json.Unmarshal([]byte(cachedData), &account); err == nil {
			return account, nil
		}
	}

	query := fmt.Sprintf("SELECT id, currency_code, active_balance, frozen_balance, user_id FROM %s WHERE id=$1", accountTable)
	err = r.db.Get(&account, query, id)

	if err == nil {
		accountData, _ := json.Marshal(account)
		r.cache.Set(cacheKey, string(accountData))
	}

	return account, err
}

func (r *AccountPostgres) GetAccountsByUserId(userId int) ([]entity.Account, error) {
	var accounts []entity.Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=%d", accountTable, userId)
	err := r.db.Select(&accounts, query)

	if err == nil {
		for _, account := range accounts {
			cacheKey := fmt.Sprintf("account_%d", account.Id)
			accountData, _ := json.Marshal(account)
			r.cache.Set(cacheKey, string(accountData))
		}
	}

	return accounts, err
}

func (r *AccountPostgres) UpdateAccount(id int, account entity.Account) error {
	query := fmt.Sprintf("UPDATE %s SET currency_code=$1, active_balance=$2, frozen_balance=$3 WHERE id=$4", accountTable)
	_, err := r.db.Exec(query, account.CurrencyCode, account.ActiveBalance, account.FrozenBalance, id)

	cacheKey := fmt.Sprintf("account_%d", id)

	if err == nil {
		accountData, _ := json.Marshal(account)
		r.cache.Set(cacheKey, string(accountData))
	}

	return err
}
