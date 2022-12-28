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

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
		return
	}
	dbQueries = db.New(conn)
	os.Exit(m.Run())
}
