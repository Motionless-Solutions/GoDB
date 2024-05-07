package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"GoDB/sqlc" // Adjust this import path according to your project structure

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://maalem:secretugtk@localhost:5432/Motionless?sslmode=disable"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = sqlc.New(testDB)

	// Run tests
	exitCode := m.Run()

	// Close the database connection after tests are finished
	if err := testDB.Close(); err != nil {
		log.Fatal("error closing test database:", err)
	}

	os.Exit(exitCode)
}
