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

func NewBeerRepositoryDatabase(db *gorm.DB) BeerRepositoryDatabase {
	return BeerRepositoryDatabase{client: db}
}
