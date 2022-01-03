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

	newBeerCreated, err := serviceBeer.Create(newBeerRequest)

	if err != nil {
		t.Error("Test failed while creating error for new beer")
	}

	if newBeerCreated.Id != beerCreated.Id {
		t.Error("Test failed when matching new beer")
	}
}
