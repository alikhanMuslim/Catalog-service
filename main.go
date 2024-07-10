package main

import (
	"database/sql"
	"log"

	"github.com/alikhanMuslim/Catalog-service/api"
	db "github.com/alikhanMuslim/Catalog-service/db/sqlc"
	"github.com/alikhanMuslim/Catalog-service/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	conn, err := sql.Open(config.DriverName, config.DriveSource)

	if err != nil {
		log.Fatal("cannot open database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(*store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("Cannot start the server")
	}
}
