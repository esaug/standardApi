package main

import (
	"database/sql"
	"log"
	"standardApi/cmd/api"
	"standardApi/config"
	"standardApi/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLIntegration(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.PublicHost,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
