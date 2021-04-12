package handlers

import (
	"net/http"
	"os"

	"github.com/DalinarLG/blog/middlewares"
	"github.com/DalinarLG/blog/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	r := mux.NewRouter()

	r.HandleFunc("/insertuser", middlewares.CheckDB(routers.InsertUser)).Methods("POST")
	r.HandleFunc("/loginuser", middlewares.CheckDB(routers.LoginUser)).Methods("POST")
	r.HandleFunc("/updateuser", middlewares.CheckDB(middlewares.ValidateToken(routers.UpdateUser))).Methods("PUT")
	r.HandleFunc("/uploadavatar", middlewares.CheckDB(middlewares.ValidateToken(routers.UploadAvatar))).Methods("PUT")
	r.HandleFunc("/insertcategory", middlewares.CheckDB(middlewares.ValidateToken(routers.InsertCategory))).Methods("POST")
	r.HandleFunc("/deletecategory", middlewares.CheckDB(middlewares.ValidateToken(routers.DeleteCategory))).Methods("DELETE")
	r.HandleFunc("/getcategories", middlewares.CheckDB(middlewares.ValidateToken(routers.GetCategories))).Methods("GET")
	r.HandleFunc("/updatecategory", middlewares.CheckDB(middlewares.ValidateToken(routers.UpdateCategory))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":"+PORT, handler)
}
