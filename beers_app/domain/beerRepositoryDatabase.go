package domain

import (
	"log"

	"gorm.io/gorm"
)

type BeerRepositoryDatabase struct {
	client *gorm.DB
}

func (brd BeerRepositoryDatabase) FindAll() ([]Beer, error) {
	var beers []Beer
	resultSet := brd.client.Find(&beers)
	if resultSet.Error != nil {
		log.Fatalln("Error has been encountered: ", resultSet.Error)
		return nil, resultSet.Error
	}
	return beers, nil
}

func (brd BeerRepositoryDatabase) FindOne(beer_id int) (*Beer, error) {
	var beer Beer
	result := brd.client.First(&beer, "id = ?", beer_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &beer, nil
}

func NewBeerRepositoryDatabase(db *gorm.DB) BeerRepositoryDatabase {
	return BeerRepositoryDatabase{client: db}
}
