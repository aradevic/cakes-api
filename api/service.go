package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"alexco.waracle.com/cakes/data"
)

func GetCakes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) > 0 {
		fmt.Fprintln(w, data.FindCakes(q))
	} else {
		fmt.Fprintln(w, data.GetCakes())
	}
}

func AddCake(w http.ResponseWriter, r *http.Request) {

	cake := parseJSON(r)
	data.AddCake(cake)
}

func DeleteCakes(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	data.DeleteById(id)
}

func FindCakesById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, data.FindById(id))
}

func parseJSON(r *http.Request) data.Cake {
	var cake data.Cake
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&cake)
	if err != nil {
		panic(err)
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		panic(err)
	}
	return cake
}
