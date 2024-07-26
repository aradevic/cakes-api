package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"alexco.waracle.com/cakes/api"
	"alexco.waracle.com/cakes/repo"
)

var s api.Service

func init() {
	var err error
	var db *sql.DB
	db, err = sql.Open("pgx", "host=localhost port=5432 user=user password=hellochat dbname=cakes")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Db!")

	s = api.Service{
		D: &repo.PostresDBRepo{Db: db},
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cakes", s.GetCakes)
	mux.HandleFunc("POST /cakes", s.AddCake)
	mux.HandleFunc("DELETE /cakes/{id}", s.DeleteCakes)
	mux.HandleFunc("GET /cakes/{id}", s.FindCakesById)

	fmt.Println("Connected to server!")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
