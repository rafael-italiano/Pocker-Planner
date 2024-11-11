package service

import (
	"database/sql"
)

type entryType int

const (
	Debit entryType = iota
	Credit
)

type Entry struct {
	ID            int
	TransactionID int
	AccountID     int
	Value         float64
}

type EntryService struct {
	db *sql.DB
}

func NewEntryService(db *sql.DB) *EntryService {
	return &EntryService{db: db}
}

func (s *EntryService) CreateEntry(entry *Entry) error {
	query := "INSERT INTO entries (transaction_id, account_id, value) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, entry.TransactionID, entry.AccountID, entry.Value)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	entry.ID = int(lastInsertID)
	return nil
}

func (s *EntryService) GetEntries() ([]Entry, error) {
	query := "SELECT id, transaction_id, account_id, value FROM entries"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.TransactionID, &entry.AccountID, &entry.Value)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (s *EntryService) GetEntryByID(id int) (*Entry, error) {
	query := "SELECT id, transaction_id, account_id, value FROM entries WHERE id = ?"
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var entry Entry
	err = row.Scan(&entry.ID, &entry.TransactionID, &entry.AccountID, &entry.Value)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (s *EntryService) GetEntryByTransactionID(transaction_id int) ([]Entry, error) {
	query := "SELECT id, transaction_id, account_id, value FROM entries WHERE transaction_id = ?"
	rows, err := s.db.Query(query, transaction_id)
	if err != nil {
		return nil, err
	}
	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.TransactionID, &entry.AccountID, &entry.Value)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (s *EntryService) GetEntryByAccountID(account_id int) ([]Entry, error) {
	query := "SELECT id, transaction_id, account_id, value FROM entries WHERE transaction_id = ?"
	rows, err := s.db.Query(query, account_id)
	if err != nil {
		return nil, err
	}
	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.TransactionID, &entry.AccountID, &entry.Value)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
