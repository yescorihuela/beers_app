package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yescorihuela/beers_app/api"
	realDomain "github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/errs"
	mocks "github.com/yescorihuela/beers_app/mocks/domain"
)

func TestBeerServiceCreate(t *testing.T) {
	newBeerRequest := api.NewBeerRequest{
		Id:       100,
		Name:     "",
		Brewery:  "Kunstmann",
		Country:  "Chile",
		Price:    0,
		Currency: "CLP",
	}

	beerService := NewBeerService(nil)
	_, err := beerService.Create(newBeerRequest)

	if err == nil {
		t.Error("Failed while testing the new beer validation")
	}
}

func TestBeerServiceCannotCreateNewBeer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	newBeerRequest := api.NewBeerRequest{
		Id:       100,
		Name:     "Torobayo",
		Brewery:  "Kunstmann",
		Country:  "Chile",
		Price:    1000,
		Currency: "CLP",
	}

	newBeer := realDomain.Beer{
		Id:       newBeerRequest.Id,
		Name:     newBeerRequest.Name,
		Brewery:  newBeerRequest.Brewery,
		Country:  newBeerRequest.Country,
		Price:    newBeerRequest.Price,
		Currency: newBeerRequest.Currency,
	}

	mockBeerRepo.EXPECT().Create(newBeer).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	_, err := serviceBeer.Create(newBeerRequest)

	if err == nil {
		t.Error("Test failed while validating error for new beer")
	}

}

func TestCreateNewBeerSuccessfully(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	newBeerRequest := api.NewBeerRequest{
		Id:       100,
		Name:     "Torobayo",
		Brewery:  "Kunstmann",
		Country:  "Chile",
		Price:    1000,
		Currency: "CLP",
	}

	newBeer := realDomain.Beer{
		Id:       newBeerRequest.Id,
		Name:     newBeerRequest.Name,
		Brewery:  newBeerRequest.Brewery,
		Country:  newBeerRequest.Country,
		Price:    newBeerRequest.Price,
		Currency: newBeerRequest.Currency,
	}

	beerCreated := newBeer
	// errs.NewUnexpectedError("Unexpected database error")
	mockBeerRepo.EXPECT().Create(newBeer).Return(&beerCreated, nil)
	newBeerCreated, err := serviceBeer.Create(newBeerRequest)

	if err != nil {
		t.Error("Test failed while creating error for new beer")
	}

	if newBeerCreated.Id != beerCreated.Id {
		t.Error("Test failed when matching new beer")
	}
}

func TestGetBeerService(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	beerID := 100

	currentBeer := realDomain.Beer{
		Id:       100,
		Name:     "Torobayo",
		Brewery:  "Kunstmann",
		Country:  "Chile",
		Price:    1000,
		Currency: "CLP",
	}

	mockBeerRepo.EXPECT().FindOne(beerID).Return(&currentBeer, nil)
	beerRequested, _ := serviceBeer.GetBeer(beerID)

	if beerRequested.Id != currentBeer.Id {
		t.Error("Test failed when matching existing beer")
	}
}

func TestGetBeerServiceNotFound(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	beerID := 2022

	mockBeerRepo.EXPECT().FindOne(beerID).Return(nil, errs.NewNotFoundError("Not found"))
	_, err := serviceBeer.GetBeer(beerID)

	if err == nil {
		t.Error("Test failed when searching not existing beer")
	}
}

func TestGetAllBeersService(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	allBeers := []realDomain.Beer{
		{
			Id:       100,
			Name:     "Torobayo",
			Brewery:  "Kunstmann",
			Country:  "Chile",
			Price:    1000,
			Currency: "CLP",
		},
		{
			Id:       101,
			Name:     "Kristal Zero",
			Brewery:  "Kristal",
			Country:  "Chile",
			Price:    500,
			Currency: "CLP",
		},
	}

	mockBeerRepo.EXPECT().FindAll().Return(allBeers, nil)
	beersRequested, _ := serviceBeer.GetAllBeers()

	if len(allBeers) != len(beersRequested) {
		t.Error("Test failed when matching all existing beers")
	}
}

func TestGetAllBeersEmptyService(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockBeerRepo := mocks.NewMockBeerRepository(controller)
	serviceBeer := NewBeerService(mockBeerRepo)

	mockBeerRepo.EXPECT().FindAll().Return(nil, errs.NewNotFoundError("Empty model"))
	beersRequested, _ := serviceBeer.GetAllBeers()

	if len(beersRequested) != 0 {
		t.Error("Test failed when searching for empty model")
	}
}
