package service

import (
	"BWG/entity"
	"BWG/internal/repo"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	GetUser(id int) (entity.User, error)
	GetUsers(options entity.Options) ([]entity.User, error)
}

type Account interface {
	CreateAccount(account entity.Account) (int, error)
	GetAccount(accountId int) (entity.Account, error)
	GetAccountsByUserId(userId int) ([]entity.Account, error)
	//UpdateAccount()
}

type Transaction interface {
	GetTransaction(id int) (entity.Transaction, error)
	GetTransactions(options entity.Options) ([]entity.Transaction, error)
	GetTransactionsByAccountId(accountId int) ([]entity.Transaction, error)
	Invoice(invoice entity.Invoice) (int, error)
	Withdraw(withdraw entity.Withdraw) (int, error)
	InitTransaction(id int) (string, error)
}

type Service struct {
	User
	Account
	Transaction
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		User:        NewUserService(repo.User),
		Account:     NewAccountService(repo.Account, repo.User),
		Transaction: NewTransactionService(repo.Transaction, repo.Account),
	}
}
