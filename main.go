package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DataBase struct {
	*sql.DB
}

func NewDataBase() (*DataBase, error) {
	connStr := "user=postgres dbname=postgres sslmode=disable password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("successfuly database connected!")
	return &DataBase{
		db,
	}, nil
}

func main() {
	db, err := NewDataBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
