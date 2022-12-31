package main

import (
	"database/sql"
	"log"

	"github.com/atlast999/project3be/api"
	"github.com/atlast999/project3be/db/transaction"
	"github.com/atlast999/project3be/helper"
	_ "github.com/lib/pq"
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
	server, err := api.NewServer(config, txInstance)
	if err != nil {
		log.Fatal("Cannot initiate server: ", err)
	}
	err = server.StartServer()
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
