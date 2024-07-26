package repo

import (
	"net/url"

	"alexco.waracle.com/cakes/model"
)

type TestDBRepo struct{}

func (t *TestDBRepo) GetCakes() (cakes []model.Cake) {
	return nil
}
func (t *TestDBRepo) AddCake(cake model.Cake) (id int) {
	return 0
}
func (t *TestDBRepo) FindCakes(v url.Values) (cakes []model.Cake) {
	return nil
}
func (t *TestDBRepo) FindById(id int) (cake model.Cake) {
	return model.Cake{
		Id:        1,
		Name:      "lemon chesse cake",
		Comment:   "the best",
		YumFactor: 5,
	}
}
func (t *TestDBRepo) DeleteById(id int) {
}
