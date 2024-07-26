package repo

import (
	"net/url"

	"alexco.waracle.com/cakes/model"
)

type DatabaseRepo interface {
	GetCakes() (cakes []model.Cake)
	AddCake(cake model.Cake) (id int)
	FindCakes(v url.Values) (cakes []model.Cake)
	FindById(id int) (cake model.Cake)
	DeleteById(id int)
}
