package main

import (
	"database/sql"
	"log"
	"picsaver/server/api"
	"picsaver/server/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/picsaverdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main()  {
    //ladda in configs
    //db settings

    conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("Cannot connect to db:", err)
    }

    store := db.NewStore(conn)
    server := api.NewServer(store)

    //Fixa address config
    err = server.Start(serverAddress)
    if err != nil {
        log.Fatal("Cannot start server:", err) 
    }
}
