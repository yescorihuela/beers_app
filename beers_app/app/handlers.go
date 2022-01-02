package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/services"
)

type BeerHandlers struct {
	serviceBeer     services.BeerService
	serviceCurrency services.CurrencyService
}

var response api.DescriptionResponse

const (
	successMessage = "Operación exitosa"
	beerCreated    = "Cerveza creada"
)

func (bh *BeerHandlers) GetAllBeers(ctx *gin.Context) {
	beers, err := bh.serviceBeer.GetAllBeers()
	if err != nil {
		response.NewDescriptionResponse(err.Message, nil)
		ctx.JSON(err.Code, response)
		return
	}

	response.NewDescriptionResponse(successMessage, beers)
	ctx.JSON(http.StatusOK, response)
}

func (bh *BeerHandlers) GetBeer(ctx *gin.Context) {
	beerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.NewDescriptionResponse(err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	beer, serviceError := bh.serviceBeer.GetBeer(beerId)
	if serviceError != nil {
		response.NewDescriptionResponse(serviceError.Message, nil)
		ctx.JSON(serviceError.Code, response)
		return
	}
	response.NewDescriptionResponse(successMessage, beer)
	ctx.JSON(http.StatusOK, response)
}

func (bh *BeerHandlers) GetBeerByBox(ctx *gin.Context) {
	beerId, beerIdErr := strconv.Atoi(ctx.Param("id"))
	toCurrency := ctx.Query("currency")
	beerQuantity, beerQuantityErr := strconv.ParseFloat(ctx.DefaultQuery("quantity", "6"), 32)

	if beerIdErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, beerIdErr)
		return
	}

	if toCurrency == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Currency is required"})
		return
	}

	if beerQuantityErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, beerQuantityErr)
		return
	}

	beer, serviceError := bh.serviceBeer.GetBeerByBox(bh.serviceCurrency, beerId, float32(beerQuantity), toCurrency)
	if serviceError != nil {
		response.NewDescriptionResponse(serviceError.Message, nil)
		ctx.JSON(serviceError.Code, response)
		return
	}
	response.NewDescriptionResponse(successMessage, beer)
	ctx.JSON(http.StatusOK, response)
}

func (bh *BeerHandlers) Create(ctx *gin.Context) {
	var newBeer api.NewBeerRequest
	err := ctx.BindJSON(&newBeer)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	_, serviceError := bh.serviceBeer.Create(newBeer)
	if serviceError != nil {
		response.NewDescriptionResponse(serviceError.Message, nil)
		ctx.JSON(serviceError.Code, response)
		return
	}
	response.NewDescriptionResponse(beerCreated, nil)
	ctx.JSON(http.StatusCreated, response)
}
