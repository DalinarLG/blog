package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB = getdb()

func getdb() *sql.DB {
	dbuser := "root"
	dbpass := ""
	dbName := "blog"

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Database connected...")
	return db
}

func CheckDB() bool {
	err := DB.Ping()
	if err != nil {
		return false
	}
	return err == nil
}
