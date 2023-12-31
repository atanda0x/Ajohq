package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/atanda0x/Ajohq/db/sqlc"
	"github.com/atanda0x/Ajohq/util"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
