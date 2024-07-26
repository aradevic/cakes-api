package model

import (
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Cake struct {
	Id        int
	Name      string
	Comment   string
	ImageUrl  string
	YumFactor int
}
