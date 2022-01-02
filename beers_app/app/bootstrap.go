package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/yescorihuela/beers_app/config"
	"github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/services"
	"gorm.io/gorm"
)

const timeout = 5 * time.Second

type Server struct {
	engine      *gin.Engine
	httpAddr    string
	database    *gorm.DB
	httpClient  *http.Client
	redisClient *redis.Client
}

func NewServer(host string, port uint) Server {
	server := Server{
		engine:      gin.Default(), // New if your need incorporate middleware or your own logger | Default is better this case
		httpAddr:    fmt.Sprintf("%s:%d", host, port),
		database:    config.ConnectDatabase(),
		httpClient:  &http.Client{Timeout: timeout},
		redisClient: config.ConnectoRedis(),
	}
	server.registerRoutes()
	return server
}

func NewRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return rdb
}

func (s *Server) Run() error {
	log.Println("Server running on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {

	// bh := BeerHandlers{service: services.NewBeerService(domain.NewBeerRepositoryStub())}
	bh := BeerHandlers{
		serviceBeer:     services.NewBeerService(domain.NewBeerRepositoryDatabase(s.database)),
		serviceCurrency: services.NewCurrencyService(domain.NewCurrencyRepositoryExternal(s.httpClient)),
	}
	s.engine.GET("/beers", bh.GetAllBeers)
	s.engine.GET("/beers/:id", bh.GetBeer)
	s.engine.GET("/beers/:id/boxprice", bh.GetBeerByBox)
	s.engine.POST("/beers/", bh.Create)
}
