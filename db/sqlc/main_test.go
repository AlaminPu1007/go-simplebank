/*
	Wrote a unit test to test db connection
	This is main test file where db will be connect.
*/

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alaminpu1007/simplebank/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

var testingQueries *Queries
var testDB *sql.DB

func TestMain(t *testing.M) {

	// load from app.env
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Count not load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("DB connection is not possible")
	}

	testingQueries = New(testDB)

	os.Exit(t.Run())
}
