package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alikhanMuslim/Catalog-service/utils"
	_ "github.com/lib/pq"
)

var testqueries *Queries

func TestMain(m *testing.M) {

	if err := SetupTestDb(); err != nil {
		log.Fatalf("error connect to database %v", err)
	}

	os.Exit(m.Run())
}

func SetupTestDb() error {

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(config.DriverName, config.DriveSource)

	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	testqueries = New(conn)

	return nil

}
