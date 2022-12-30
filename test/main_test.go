package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/db/transaction"
	"github.com/atlast999/project3be/helper"
	_ "github.com/lib/pq"
)

var dbQueries *db.Queries
var dbInstance *sql.DB
var txInstance *transaction.TxInstance

func TestMain(m *testing.M) {
	var err error
	config, err := helper.LoadConfig("../.")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}
	dbInstance, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
		return
	}
	dbQueries = db.New(dbInstance)
	txInstance = transaction.NewTxInstance(dbInstance)
	os.Exit(m.Run())
}
