package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

type DB struct {
	db *sql.DB
}

var postgres *DB

func NewDB(driver, psqlInfo string) *DB {
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &DB{db: db}
}

func InsertVote(w http.ResponseWriter, r *http.Request) {
	browser := mux.Vars(r)
	sqlStatement := "UPDATE votes SET " + browser["id"] + " = " + browser["id"] + " + 1;"
	res, err := postgres.db.Exec(sqlStatement)
	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}
	fmt.Println("New vote for :", browser, "result :", res)
	w.WriteHeader(200)
}

func main() {
	postgres = NewDB("postgres", psqlInfo)
	router := mux.NewRouter()
	router.HandleFunc("/insert_vote/{id}", InsertVote).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
