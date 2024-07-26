package repo

import (
	"database/sql"
	"net/url"

	"alexco.waracle.com/cakes/model"
)

type PostresDBRepo struct {
	Db *sql.DB
}

func (d *PostresDBRepo) GetCakes() (cakes []model.Cake) {
	rows, err := d.Db.Query("SELECT id, name, comment, url, yum FROM cakes")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cake := model.Cake{}
		if err = rows.Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
			panic(err)
		}
		cakes = append(cakes, cake)
	}
	rows.Close()
	return
}

func (d *PostresDBRepo) AddCake(cake model.Cake) (id int) {
	statement := "INSERT INTO cakes (id, name, comment, url, yum) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := d.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	if err := stmt.QueryRow(cake.Id, cake.Name, cake.Comment, cake.ImageUrl, cake.YumFactor).Scan(&id); err != nil {
		panic(err)
	}
	return
}

func (d *PostresDBRepo) FindCakes(v url.Values) (cakes []model.Cake) {
	statement := "SELECT id, name, comment, url, yum FROM cakes"
	var findBy string
	if v.Has("yum") {
		findBy = "yum"
		statement = statement + " " + "WHERE yum <= $1"
	} else if v.Has("name") {
		findBy = "name"
		statement = statement + " " + "WHERE name = $1"
	}
	stmt, err := d.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(v.Get(findBy))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cake := model.Cake{}
		if err := rows.Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
			panic(err)
		}
		cakes = append(cakes, cake)
	}

	return
}

func (d *PostresDBRepo) FindById(id int) (cake model.Cake) {

	statement := "SELECT id, name, comment, url, yum FROM cakes WHERE id = $1"
	stmt, err := d.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	if err := stmt.QueryRow(id).Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
		panic(err)
	}
	return
}

func (d *PostresDBRepo) DeleteById(id int) {
	statement := "DELETE FROM cakes where id = $1"
	stmt, err := d.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
