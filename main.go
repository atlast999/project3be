package main

import (
	"database/sql"
	"log"

	"github.com/atlast999/project3be/api"
	"github.com/atlast999/project3be/db/transaction"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:12345678@project-db.cnoos3wsb1s6.ap-northeast-1.rds.amazonaws.com/appdb"
)

func main() {
	dbInstance, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	txInstance := transaction.NewTxInstance(dbInstance)
	server := api.NewServer(txInstance)
	err = server.StartServer()
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
