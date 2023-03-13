package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Sqlite struct {
	Db *sql.DB
}

func NewSqlite() (*Sqlite, error) {
	db, err := sql.Open("sqlite3", "quotation.db")
	if err != nil {
		return &Sqlite{}, fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS quotation (id INTEGER PRIMARY KEY, value DECIMAL(5,15))")
	if err != nil {
		return &Sqlite{}, fmt.Errorf("unable to create database")
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) AddQuotation() (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO quotation(value) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("error inserting data: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("error inserting data: %v", err)
	}

	id, _ := res.LastInsertId()

	return id, nil
}
