package main

import (
	"fmt"
	"go-xsis-movie/controllers"
	"go-xsis-movie/helpers"
	"go-xsis-movie/models/db"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helpers.Logger("error", "Error getting env")
	}

	db.Init()

	ctm := mux.NewRouter()

	ctm.HandleFunc("/Movies", controllers.C_AddMovie).Methods("POST")
	ctm.HandleFunc("/Movies", controllers.C_GetAllMovie).Methods("GET")
	ctm.HandleFunc("/Movies/{ID}", controllers.C_GetSingleMovieId).Methods("GET")
	ctm.HandleFunc("/Movies/{ID}", controllers.C_UpdateSingleMovieId).Methods("PATCH")
	ctm.HandleFunc("/Movies/{ID}", controllers.C_DeleteMovie).Methods("DELETE")

	p := os.Getenv("PORT")
	h := ctm
	s := new(http.Server)
	s.Handler = h
	s.Addr = ":" + p
	fmt.Println("Server run in Port ", s.Addr)
	s.ListenAndServe()
}
