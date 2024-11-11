package service

import "database/sql"

type Account struct {
	ID         int
	ParentCode int
	Code       int
	Name       string
}

type AccountService struct {
	db *sql.DB
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateAccount(account *Account) error {
	query := "INSERT INTO accounts (id, code) VALUES (?, ?)"
	result, err := s.db.Exec(query, account.ID, account.Code)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(lastInsertID)

	return nil
}

func (s *AccountService) GetAccountsByParentCode(account *Account) error {
	query := "INSERT INTO accounts (id, parent_code, code) VALUES (?, ?)"
	result, err := s.db.Exec(query, account.ID, account.Code)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(lastInsertID)

	return nil
}
