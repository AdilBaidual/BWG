package repo

import (
	"BWG/entity"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db    *sqlx.DB
	cache *RedisRepository
}

func NewTransactionPostgres(db *sqlx.DB, cache *RedisRepository) *TransactionPostgres {
	return &TransactionPostgres{db: db, cache: cache}
}

func (r *TransactionPostgres) CreateTransaction(transaction entity.Transaction) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (currency_code, transaction_status, sender_account_id, recipient_account_id, amount)"+
		" values ($1, $2, $3, $4, $5) RETURNING id", transactionTable)
	row := r.db.QueryRow(query, transaction.CurrencyCode, transaction.TransactionStatus, transaction.SenderAccountID, transaction.RecipientAccountID, transaction.Amount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TransactionPostgres) GetTransaction(id int) (entity.Transaction, error) {
	var transaction entity.Transaction

	cacheKey := fmt.Sprintf("transaction_%d", id)
	cachedData, err := r.cache.Get(cacheKey)
	if err == nil {
		if err := json.Unmarshal([]byte(cachedData), &transaction); err == nil {
			return transaction, nil
		}
	}

	query := fmt.Sprintf("SELECT id, currency_code, transaction_status, sender_account_id, recipient_account_id, amount, transaction_date FROM %s WHERE id=$1", transactionTable)
	err = r.db.Get(&transaction, query, id)

	if err == nil {
		TransactionData, _ := json.Marshal(transaction)
		r.cache.Set(cacheKey, string(TransactionData))
	}

	return transaction, err
}

func (r *TransactionPostgres) GetTransactions(query string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Select(&transactions, query)

	if err == nil {
		for _, transaction := range transactions {
			cacheKey := fmt.Sprintf("transaction_%d", transaction.Id)
			transactionData, _ := json.Marshal(transaction)
			r.cache.Set(cacheKey, string(transactionData))
		}
	}

	return transactions, err
}

func (r *TransactionPostgres) GetTransactionsByAccountId(accountId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	query := fmt.Sprintf("SELECT * FROM %s WHERE sender_account_id=%d OR recipient_account_id=%d", transactionTable, accountId, accountId)
	err := r.db.Select(&transactions, query)

	if err == nil {
		for _, transaction := range transactions {
			cacheKey := fmt.Sprintf("transaction_%d", transaction.Id)
			transactionData, _ := json.Marshal(transaction)
			r.cache.Set(cacheKey, string(transactionData))
		}
	}

	return transactions, err
}

func (r *TransactionPostgres) UpdateTransaction(id int, transaction entity.Transaction) error {
	query := fmt.Sprintf("UPDATE %s SET transaction_status=$1, transaction_date=$2 WHERE id=$3", transactionTable)
	_, err := r.db.Exec(query, transaction.TransactionStatus, transaction.TransactionDate, id)

	cacheKey := fmt.Sprintf("transaction_%d", id)

	if err == nil {
		transactionData, _ := json.Marshal(transaction)
		r.cache.Set(cacheKey, string(transactionData))
	}

	return err
}
