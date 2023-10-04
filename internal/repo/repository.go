package repo

import (
	"BWG/entity"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	GetUser(id int) (entity.User, error)
	GetUsers(query string) ([]entity.User, error)
}

type Account interface {
	CreateAccount(account entity.Account) (int, error)
	GetAccount(id int) (entity.Account, error)
	GetAccountsByUserId(userId int) ([]entity.Account, error)
	UpdateAccount(id int, account entity.Account) error
}

type Transaction interface {
	CreateTransaction(transaction entity.Transaction) (int, error)
	GetTransaction(id int) (entity.Transaction, error)
	GetTransactions(query string) ([]entity.Transaction, error)
	GetTransactionsByAccountId(accountId int) ([]entity.Transaction, error)
	UpdateTransaction(id int, transaction entity.Transaction) error
}

type Repository struct {
	User
	Account
	Transaction
}

func NewRepo(db *sqlx.DB, client *redis.Client) *Repository {
	cache := NewRedisRepository(client)
	return &Repository{
		User:        NewUserPostgres(db, cache),
		Account:     NewAccountPostgres(db, cache),
		Transaction: NewTransactionPostgres(db, cache),
	}
}
