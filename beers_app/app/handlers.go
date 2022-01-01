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
	beers, _ := bh.service.GetAllBeers()
	ctx.JSON(http.StatusOK, gin.H{"beers": beers})
}

func (bh *BeerHandlers) GetBeer(ctx *gin.Context) {
	beer_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	beer, err := bh.service.GetBeer(beer_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{"beer": beer})
}

func (bh *BeerHandlers) Create(ctx *gin.Context) {
	var newBeer api.NewBeerRequest
	err := ctx.BindJSON(&newBeer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	beer, serviceError := bh.service.Create(newBeer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.AsMessage()})
	}
	ctx.JSON(http.StatusOK, beer)
}
