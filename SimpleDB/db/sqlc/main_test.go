package db

import (
	"database/sql"
	"os"
	"simplebank/db/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		panic(err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic(err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
