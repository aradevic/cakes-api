package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"alexco.waracle.com/cakes/model"
	"alexco.waracle.com/cakes/repo"
)

type Service struct {
	D repo.DatabaseRepo
}

func (s *Service) GetCakes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) > 0 {
		fmt.Fprintln(w, s.D.FindCakes(q))
	} else {
		fmt.Fprintln(w, s.D.GetCakes())
	}
}

func (s *Service) AddCake(w http.ResponseWriter, r *http.Request) {

	cake := parseJSON(r)
	s.D.AddCake(cake)
}

func (s *Service) DeleteCakes(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	s.D.DeleteById(id)
}

func (s *Service) FindCakesById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, s.D.FindById(id))
}

func parseJSON(r *http.Request) model.Cake {
	var cake model.Cake
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
