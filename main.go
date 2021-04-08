package main

import (
	"log"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/handlers"
)


func main() {
	if !db.CheckDB(){
		log.Fatal("No database connection")
		return
	}

	handlers.Handlers()
}
