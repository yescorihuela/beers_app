package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/services"
)

type BeerHandlers struct {
	service services.BeerService
}

var response api.DescriptionResponse

const (
	successMessage = "Operación exitosa"
	beerCreated    = "Cerveza creada"
)

func (bh *BeerHandlers) GetAllBeers(ctx *gin.Context) {
	beers, err := bh.service.GetAllBeers()
	if err != nil {
		response.NewDescriptionResponse(err.Message, nil)
		ctx.JSON(err.Code, response)
		return
	}

	response.NewDescriptionResponse(successMessage, beers)
	ctx.JSON(http.StatusOK, response)
}

func (bh *BeerHandlers) GetBeer(ctx *gin.Context) {
	beer_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.NewDescriptionResponse(err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	beer, serviceError := bh.service.GetBeer(beer_id)
	if serviceError != nil {
		response.NewDescriptionResponse(serviceError.Message, nil)
		ctx.JSON(serviceError.Code, response)
		return
	}
	response.NewDescriptionResponse(successMessage, beer)
	ctx.JSON(http.StatusOK, response)
}

func (bh *BeerHandlers) GetBeerByBox(ctx *gin.Context) {
	beer_id, beer_id_err := strconv.Atoi(ctx.Param("id"))
	beer_currency := ctx.Query("currency")
	beer_quantity, beer_quantity_err := strconv.ParseFloat(ctx.DefaultQuery("quantity", "6"), 32)

	if beer_id_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, beer_id_err)
		return
	}

	if beer_currency == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Currency is required"})
		return
	}

	if beer_quantity_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, beer_quantity_err)
		return
	}

	beer, serviceError := bh.service.GetBeerByBox(beer_id, float32(beer_quantity), beer_currency)
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
	_, serviceError := bh.service.Create(newBeer)
	if serviceError != nil {
		response.NewDescriptionResponse(serviceError.Message, nil)
		ctx.JSON(serviceError.Code, response)
		return
	}
	response.NewDescriptionResponse(beerCreated, nil)
	ctx.JSON(http.StatusCreated, response)
}
