package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/beers_app/services"
)

type BeerHandlers struct {
	service services.BeerService
}

func (bh *BeerHandlers) GetAllBeers(ctx *gin.Context) {
	beers, _ := bh.service.GetAllBeers()
	ctx.JSON(http.StatusOK, gin.H{"beers": beers})
}

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
