package app

import (
	"bytes"
	"encoding/json"
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

func TestCreateNewBeerHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockBeerService(ctrl)
	dummyBeer := api.NewBeerRequest{Id: 10, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}
	dummyResponse := api.BeerResponse{Id: 10, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}

	mockService.EXPECT().Create(dummyBeer).Return(&dummyResponse, nil)
	bh := BeerHandlers{serviceBeer: mockService}
	r := gin.Default()
	r.POST("/beers/", bh.Create)
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&dummyBeer)
	request, _ := http.NewRequest(http.MethodPost, "/beers/", payload)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func TestCreateNewBeerThatExistsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockBeerService(ctrl)
	dummyBeer := api.NewBeerRequest{Id: 10, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}
	// dummyResponse := api.BeerResponse{Id: 10, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}
	dummyError := errs.NewConflictError("El ID de la cerveza ya existe")
	mockService.EXPECT().Create(dummyBeer).Return(nil, dummyError)
	bh := BeerHandlers{serviceBeer: mockService}
	r := gin.Default()
	r.POST("/beers/", bh.Create)
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&dummyBeer)
	request, _ := http.NewRequest(http.MethodPost, "/beers/", payload)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusConflict {
		t.Error("Failed while testing the status code")
	}
}

func TestGetBeerByBoxHandler(t *testing.T) {
	controllerBeerService := gomock.NewController(t)
	controllerCurrencyService := gomock.NewController(t)
	defer controllerBeerService.Finish()
	defer controllerCurrencyService.Finish()
	mockServiceBeer := mocks.NewMockBeerService(controllerBeerService)
	mockCurrencyService := mocks.NewMockCurrencyService(controllerCurrencyService)

	// dummyBeer := api.BeerResponse{Id: 1, Name: "Torobayo", Brewery: "Kunstmann", Country: "Chile", Price: 1200, Currency: "CLP"}
	dummyPrice := api.BeerBoxTotalPrice{
		Price: 10000,
	}
	beerDummyParam := 1
	quantityDummyParam := float32(6.0)
	currencyDummyParam := "EUR"

	mockServiceBeer.EXPECT().GetBeerByBox(mockCurrencyService, beerDummyParam, quantityDummyParam, currencyDummyParam).Return(&dummyPrice, nil)
	bh := BeerHandlers{serviceBeer: mockServiceBeer, serviceCurrency: mockCurrencyService}
	r := gin.Default()
	r.GET("/beers/:id/boxprice", bh.GetBeerByBox)
	url := fmt.Sprintf("/beers/%d/boxprice?currency=EUR&quantity=6", beerDummyParam)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}
