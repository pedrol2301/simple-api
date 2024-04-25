package main

import (
	"simple-api/cmd/api"
	"simple-api/config"
	"simple-api/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMYSQLStoage(mysql.Config{
		User:              config.Envs.DBUser,
		Passwd:            config.Envs.DBPass,
		Net:               "tcp",
		Addr:              config.Envs.DBAddr,
		DBName:            config.Envs.DBName,
		AllowOldPasswords: true,
		ParseTime:         true,
	})

	if err != nil {
		panic(err)
	}

	server := api.NewAPIServer("localhost:8080", db)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
