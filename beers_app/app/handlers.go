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

func (bh *BeerHandlers) GetAllBeers(ctx *gin.Context) {
	beers, err := bh.service.GetAllBeers()
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.AsMessage()})
		return
	}
	ctx.JSON(http.StatusOK, beers)
}

func (bh *BeerHandlers) GetBeer(ctx *gin.Context) {
	beer_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	beer, serviceError := bh.service.GetBeer(beer_id)
	if serviceError != nil {
		ctx.JSON(serviceError.Code, serviceError.AsMessage())
		return
	}
	ctx.JSON(http.StatusOK, beer)
}

func (bh *BeerHandlers) Create(ctx *gin.Context) {
	var newBeer api.NewBeerRequest
	err := ctx.BindJSON(&newBeer)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	beer, serviceError := bh.service.Create(newBeer)
	if serviceError != nil {
		ctx.JSON(serviceError.Code, serviceError.AsMessage())
		return
	}
	ctx.JSON(http.StatusOK, beer)
}
