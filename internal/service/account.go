package service

import (
	"BWG/entity"
	"BWG/internal/repo"
)

type AccountService struct {
	repo     repo.Account
	userRepo repo.User
}

func NewAccountService(repo repo.Account, userRepo repo.User) *AccountService {
	return &AccountService{repo: repo, userRepo: userRepo}
}

func (s *AccountService) CreateAccount(account entity.Account) (int, error) {
	_, err := s.userRepo.GetUser(*account.UserId)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateAccount(account)
	return id, err
}

func (s *AccountService) GetAccount(accountId int) (entity.Account, error) {
	account, err := s.repo.GetAccount(accountId)
	return account, err
}

func (s *AccountService) GetAccountsByUserId(userId int) ([]entity.Account, error) {
	accounts, err := s.repo.GetAccountsByUserId(userId)
	return accounts, err
}

//func (s *AccountService) UpdateAccount(account entity.Account) error {
//	offset := (options.Page - 1) * options.PerPage
//	query := fmt.Sprintf("SELECT * FROM Accounts LIMIT %d OFFSET %d", options.PerPage, offset)
//	Accounts, err := s.repo.GetAll(query)
//
//	return Accounts, err
//}
