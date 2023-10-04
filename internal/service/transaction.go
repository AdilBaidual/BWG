package service

import (
	"BWG/entity"
	"BWG/internal/repo"
	"errors"
	"fmt"
)

const (
	TRANSACTION_ERROR   = "Error"
	TRANSACTION_CREATED = "Created"
	TRANSACTION_SUCCESS = "Success"
)

type TransactionService struct {
	repo        repo.Transaction
	accountRepo repo.Account
}

func NewTransactionService(repo repo.Transaction, accountRepo repo.Account) *TransactionService {
	return &TransactionService{repo: repo, accountRepo: accountRepo}
}

//func (s *TransactionService) CreateTransaction(Transaction entity.Transaction) (int, error) {
//	id, err := s.repo.CreateTransaction(Transaction)
//	return id, err
//}

func (s *TransactionService) GetTransaction(transactionId int) (entity.Transaction, error) {
	transaction, err := s.repo.GetTransaction(transactionId)
	return transaction, err
}

func (s *TransactionService) GetTransactions(options entity.Options) ([]entity.Transaction, error) {
	offset := (options.Page - 1) * options.PerPage
	query := fmt.Sprintf("SELECT * FROM transactions LIMIT %d OFFSET %d", options.PerPage, offset)
	Transactions, err := s.repo.GetTransactions(query)

	return Transactions, err
}

func (s *TransactionService) GetTransactionsByAccountId(accountId int) ([]entity.Transaction, error) {
	transactions, err := s.repo.GetTransactionsByAccountId(accountId)
	return transactions, err
}

func (s *TransactionService) Invoice(invoice entity.Invoice) (int, error) {
	var transaction entity.Transaction

	transaction.CurrencyCode = invoice.CurrencyCode
	transaction.RecipientAccountID = invoice.RecipientAccountID
	transaction.Amount = invoice.Amount
	transaction.TransactionStatus = invoice.TransactionStatus

	if invoice.TransactionStatus == TRANSACTION_ERROR {
		id, err := s.repo.CreateTransaction(transaction)
		return id, err
	}
	account, err := s.accountRepo.GetAccount(*invoice.RecipientAccountID)
	if err != nil {
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	if account.CurrencyCode != invoice.CurrencyCode {
		err = errors.New("currency code error")
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	id, err := s.repo.CreateTransaction(transaction)
	return id, err
}

func (s *TransactionService) Withdraw(withdraw entity.Withdraw) (int, error) {
	var transaction entity.Transaction

	transaction.CurrencyCode = withdraw.CurrencyCode
	transaction.SenderAccountID = withdraw.SenderAccountID
	transaction.RecipientAccountID = withdraw.RecipientAccountID
	transaction.Amount = withdraw.Amount
	transaction.TransactionStatus = withdraw.TransactionStatus

	if withdraw.TransactionStatus == TRANSACTION_ERROR {
		id, err := s.repo.CreateTransaction(transaction)
		return id, err
	}

	accountSender, err := s.accountRepo.GetAccount(*withdraw.SenderAccountID)
	if err != nil {
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	if accountSender.ActiveBalance < withdraw.Amount {
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	accountSender.ActiveBalance -= withdraw.Amount
	accountSender.FrozenBalance += withdraw.Amount

	err = s.accountRepo.UpdateAccount(accountSender.Id, accountSender)
	if err != nil {
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	var accountRecipient entity.Account
	if withdraw.RecipientAccountID != nil {
		accountRecipient, err = s.accountRepo.GetAccount(*withdraw.RecipientAccountID)
		if err != nil {
			transaction.TransactionStatus = TRANSACTION_ERROR
			id, _ := s.repo.CreateTransaction(transaction)
			return id, err
		}
	}

	if accountSender.CurrencyCode != withdraw.CurrencyCode || (withdraw.RecipientAccountID != nil && accountRecipient.CurrencyCode != withdraw.CurrencyCode) {
		err = errors.New("currency code error")
		transaction.TransactionStatus = TRANSACTION_ERROR
		id, _ := s.repo.CreateTransaction(transaction)
		return id, err
	}

	id, err := s.repo.CreateTransaction(transaction)
	return id, err
}

func (s *TransactionService) InitTransaction(id int) (string, error) {
	transaction, err := s.repo.GetTransaction(id)
	if err != nil {
		transaction.TransactionStatus = TRANSACTION_ERROR
		s.repo.UpdateTransaction(transaction.Id, transaction)
		return transaction.TransactionStatus, err
	}

	if transaction.TransactionStatus == TRANSACTION_ERROR || transaction.TransactionStatus == TRANSACTION_SUCCESS {
		err = errors.New("transaction was inited before")
		return transaction.TransactionStatus, err
	}

	accountSenderId := transaction.SenderAccountID
	accountRecipientId := transaction.RecipientAccountID

	if accountSenderId != nil {
		accountSender, err := s.accountRepo.GetAccount(*accountSenderId)
		if err != nil {
			transaction.TransactionStatus = TRANSACTION_ERROR
			s.repo.UpdateTransaction(transaction.Id, transaction)
			return transaction.TransactionStatus, err
		}

		accountSender.FrozenBalance -= transaction.Amount

		err = s.accountRepo.UpdateAccount(accountSender.Id, accountSender)
		if err != nil {
			transaction.TransactionStatus = TRANSACTION_ERROR
			s.repo.UpdateTransaction(transaction.Id, transaction)
			return transaction.TransactionStatus, err
		}
	}

	if accountRecipientId != nil {
		accountRecipient, err := s.accountRepo.GetAccount(*accountRecipientId)
		if err != nil {
			transaction.TransactionStatus = TRANSACTION_ERROR
			s.repo.UpdateTransaction(transaction.Id, transaction)
			return transaction.TransactionStatus, err
		}

		accountRecipient.ActiveBalance += transaction.Amount

		err = s.accountRepo.UpdateAccount(accountRecipient.Id, accountRecipient)
		if err != nil {
			transaction.TransactionStatus = TRANSACTION_ERROR
			s.repo.UpdateTransaction(transaction.Id, transaction)
			return transaction.TransactionStatus, err
		}
	}

	transaction.TransactionStatus = TRANSACTION_SUCCESS
	err = s.repo.UpdateTransaction(transaction.Id, transaction)

	return transaction.TransactionStatus, err
}
