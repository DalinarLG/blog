package handlers

import (
	"net/http"
	"os"

	"github.com/DalinarLG/blog/middlewares"
	"github.com/DalinarLG/blog/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)



func  Handlers(){
	r := mux.NewRouter()

	r.HandleFunc("/insertuser", middlewares.CheckDB(routers.InsertUser)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":"+PORT, handler)
}