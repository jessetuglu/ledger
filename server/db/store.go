package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Transaction err: %v, Rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func formatDbUrl(DbHost, DbPort, DbUser, DbName, DbPassword string) (string){
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
}

func ConnectToDB(username, password, url, database string) (*sql.DB, error){
	connect := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, url, database)
	conn, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("Couldn't connect to the DB")
		return nil, err
	}
	return conn, nil	
}
