package main

import (
	"fmt"
	"net/http"

	"alexco.waracle.com/cakes/api"
)

func main() {
	// var err error
	http.HandleFunc("GET /cakes", api.GetCakes)
	http.HandleFunc("POST /cakes", api.AddCake)
	http.HandleFunc("DELETE /cakes/{id}", api.DeleteCakes)
	http.HandleFunc("GET /cakes/{id}", api.FindCakesById)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("Connected to server!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
