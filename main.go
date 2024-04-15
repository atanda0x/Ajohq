package main

import (
	"database/sql"
	"log"

	"github.com/atanda0x/FintechConnect/api"
	db "github.com/atanda0x/FintechConnect/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	serverAddress = "0.0.0.0:9090"
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:ethereum@localhost:5432/fintechAPI?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
