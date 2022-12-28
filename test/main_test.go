package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/atlast999/project3be/db/gen"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgres://postgres:12345678@project-db.cnoos3wsb1s6.ap-northeast-1.rds.amazonaws.com/appdb"

var dbQueries *db.Queries
var dbInstance *sql.DB
func TestMain(m *testing.M) {
	var err error
	dbInstance, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
		return
	}
	dbQueries = db.New(dbInstance)
	os.Exit(m.Run())
}
