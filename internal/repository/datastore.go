package repository

import (
	"context"
	"database/sql"
)

type dbtx interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type DataStore interface {
	Atomic(ctx context.Context, fn func(DataStore) error) error
	User() UserRepository
	Contact() ContactRepository
}

type dataStore struct {
	conn *sql.DB
	db   dbtx
}

func NewDataStore(db *sql.DB) DataStore {
	return &dataStore{
		conn: db,
		db:   db,
	}
}

func (s *dataStore) Atomic(ctx context.Context, fn func(DataStore) error) error {
	tx, err := s.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	err = fn(&dataStore{conn: s.conn, db: tx})
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (s *dataStore) User() UserRepository {
	return NewUserRepository(s.db)
}

func (s *dataStore) Contact() ContactRepository {
	return NewContactRepository(s.db)
}
