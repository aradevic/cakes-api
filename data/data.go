package data

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Cake struct {
	Id        int
	Name      string
	Comment   string
	ImageUrl  string
	YumFactor int
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("pgx", "host=localhost port=5432 user=user password=hellochat dbname=cakes")
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Db!")
}

func GetCakes() (cakes []Cake) {
	rows, err := Db.Query("SELECT id, name, comment, url, yum FROM cakes")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cake := Cake{}
		if err = rows.Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
			panic(err)
		}
		cakes = append(cakes, cake)
	}
	rows.Close()
	return
}

func AddCake(cake Cake) (id int) {
	statement := "INSERT INTO cakes (id, name, comment, url, yum) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	if err := stmt.QueryRow(cake.Id, cake.Name, cake.Comment, cake.ImageUrl, cake.YumFactor).Scan(&id); err != nil {
		panic(err)
	}
	return
}

func FindCakes(v url.Values) (cakes []Cake) {
	statement := "SELECT id, name, comment, url, yum FROM cakes"
	var findBy string
	if v.Has("yum") {
		findBy = "yum"
		statement = statement + " " + "WHERE yum <= $1"
	} else if v.Has("name") {
		findBy = "name"
		statement = statement + " " + "WHERE name = $1"
	}
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(v.Get(findBy))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cake := Cake{}
		if err := rows.Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
			panic(err)
		}
		cakes = append(cakes, cake)
	}

	return
}

func FindById(id int) (cake Cake) {

	statement := "SELECT id, name, comment, url, yum FROM cakes WHERE id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	if err := stmt.QueryRow(id).Scan(&cake.Id, &cake.Name, &cake.Comment, &cake.ImageUrl, &cake.YumFactor); err != nil {
		panic(err)
	}
	return
}

func DeleteById(id int) {
	statement := "DELETE FROM cakes where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
