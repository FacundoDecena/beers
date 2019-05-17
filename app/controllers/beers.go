package controllers

import (
	"beers/app/models"
	"github.com/revel/revel"
)

type Beers struct {
	*revel.Controller
}

var beers = []models.Beer{
	models.Beer{0, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
	models.Beer{1, "Cocoa Psycho", "Porter", "BrewDog"},
	models.Beer{2, "American Dream", "Lager", "Mikkeller"},
}

func (b Beers) List() revel.Result{
	return b.RenderJSON(beers)
}

func (b Beers) Show(beerID int) revel.Result {

	length := len (beers)

	for i := 0; i < length; i++{
		if beerID == beers[i].ID{
			return b.RenderJSON(beers[i])
		}
	}

	return b.NotFound("Could not find beer")
}