package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/services"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
}

func NewServer(host string, port uint) Server {
	server := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	log.Println("Server running on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	bh := BeerHandlers{service: services.NewBeerService(domain.NewBeerRepositoryStub())}
	s.engine.GET("/beers", bh.GetAllBeers)
}
