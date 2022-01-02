package domain

import (
	"log"

	"github.com/yescorihuela/beers_app/errs"
	"gorm.io/gorm"
)

type BeerRepositoryDatabase struct {
	client *gorm.DB
}

func (brd BeerRepositoryDatabase) FindAll() ([]Beer, *errs.AppError) {
	var beers []Beer
	resultSet := brd.client.Find(&beers)
	if resultSet.Error != nil {
		log.Fatalln("Error has been encountered: ", resultSet.Error)
		return nil, errs.NewNotFoundError(resultSet.Error.Error())
	}
	return beers, nil
}

func (brd BeerRepositoryDatabase) FindOne(beer_id int) (*Beer, *errs.AppError) {
	var beer Beer
	result := brd.client.First(&beer, "id = ?", beer_id)
	if result.Error != nil {
		return nil, errs.NewNotFoundError("Not found")
	}
	return &beer, nil
}

func (brd BeerRepositoryDatabase) Create(beer Beer) (*Beer, *errs.AppError) {

	err := brd.FindIfExists(beer.Id)

	if err != nil {
		return nil, err
	}

	result := brd.client.Create(&beer)
	if result.Error != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	return &beer, nil
}

func (brd BeerRepositoryDatabase) FindIfExists(id uint) *errs.AppError {
	var exists bool
	var beer Beer
	subQueryExists := brd.client.Select("1").Where("id = ?", id).Model(&beer)
	resultExists := brd.client.Select("EXISTS(?) as exists", subQueryExists).Model(&beer).Find(&exists)

	if exists {
		return errs.NewConflictError("El ID de la cerveza ya existe")
	}

	if resultExists.Error != nil {
		return errs.NewUnexpectedError(resultExists.Error.Error())
	}

	if resultExists.Error != nil {
		return errs.NewUnexpectedError(resultExists.Error.Error())
	}
	return nil
}

func NewBeerRepositoryDatabase(db *gorm.DB) BeerRepositoryDatabase {
	return BeerRepositoryDatabase{client: db}
}
