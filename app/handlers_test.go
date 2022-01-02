package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/errs"
	mocks "github.com/yescorihuela/beers_app/mocks/services"
)

func TestGetAllBeersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockBeerService(ctrl)
	dummyBeers := []api.BeerResponse{
		{Id: 1, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"},
		{Id: 2, Name: "Ipa", Brewery: "Escudo", Country: "Chile", Price: 800, Currency: "CLP"},
	}
	mockService.EXPECT().GetAllBeers().Return(dummyBeers, nil)
	bh := BeerHandlers{serviceBeer: mockService}
	r := gin.Default()
	r.GET("/beers", bh.GetAllBeers)

	request, _ := http.NewRequest(http.MethodGet, "/beers", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestGetBeerHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockBeerService(ctrl)
	dummyBeer := api.BeerResponse{Id: 1, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}

	beerParam := 1

	mockService.EXPECT().GetBeer(beerParam).Return(&dummyBeer, nil)
	bh := BeerHandlers{serviceBeer: mockService}
	r := gin.Default()
	r.GET("/beers/:id", bh.GetBeer)

	request, _ := http.NewRequest(http.MethodGet, "/beers/1", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestGetBeerNotFoundHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockBeerService(ctrl)
	dummyError := errs.NewNotFoundError("Cerveza no encontrada")

	beerParam := 1000

	mockService.EXPECT().GetBeer(beerParam).Return(nil, dummyError)
	bh := BeerHandlers{serviceBeer: mockService}
	r := gin.Default()
	r.GET("/beers/:id", bh.GetBeer)

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/beers/%d", beerParam), nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}
