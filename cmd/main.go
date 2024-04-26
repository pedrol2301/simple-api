package main

import (
	"database/sql"
	"log"
	"simple-api/cmd/api"
	"simple-api/config"
	"simple-api/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMYSQLStoage(mysql.Config{
		User:              config.Envs.DBUser,
		Passwd:            config.Envs.DBPass,
		Addr:              config.Envs.DBAddr,
		DBName:            config.Envs.DBName,
		Net:               "tcp",
		AllowOldPasswords: true,
		ParseTime:         true,
	})

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	initStorage(db)
	server := api.NewAPIServer("localhost:8080", db)

	if err := server.Start(); err != nil {
		panic(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	log.Println("Connected to database")
}
