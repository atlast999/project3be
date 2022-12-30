package main

import (
	"database/sql"
	"log"

	"github.com/atlast999/project3be/api"
	"github.com/atlast999/project3be/db/transaction"
	"github.com/atlast999/project3be/helper"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:12345678@project-db.cnoos3wsb1s6.ap-northeast-1.rds.amazonaws.com/appdb"
)

func main() {
	config, err := helper.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}
	dbInstance, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	txInstance := transaction.NewTxInstance(dbInstance)
	server := api.NewServer(txInstance)
	err = server.StartServer()
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
