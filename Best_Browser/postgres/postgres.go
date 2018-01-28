package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDB(driver, psqlInfo string) (*DB, error) {
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (p *DB) GetVotes(browser string) int {
	sqlStatement := "SELECT " + browser + " FROM votes;"
	count := 0
	err := p.db.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read vote for :", browser, "is ", count)
	return count
}

func (p *DB) InsertVote(browser string) {
	sqlStatement := "UPDATE votes SET " + browser + " = " + browser + " + 1;"
	res, err := p.db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("New vote for :", browser, "result :", res)
}
