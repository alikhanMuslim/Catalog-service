package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	testqueries = New(conn)

	return nil

}
