package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

func GetVotes(w http.ResponseWriter, r *http.Request) {
	browser := mux.Vars(r)
	sqlStatement := "SELECT " + browser["id"] + " FROM votes;"
	count := 0
	err := postgres.db.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}
	fmt.Println("Read vote for :", browser, "is ", count)
	result := strconv.Itoa(count)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

func main() {
	postgres = NewDB("postgres", psqlInfo)
	router := mux.NewRouter()
	router.HandleFunc("/get_votes/{id}", GetVotes).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
