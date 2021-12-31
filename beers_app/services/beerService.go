package services

import (
	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/domain"
)

type BeerService interface {
	GetAllBeers() ([]api.BeerResponse, error)
	GetBeer(id int) (*api.BeerResponse, error)
}

type DefaultBeerService struct {
	repo domain.BeerRepository
}

func (s DefaultBeerService) GetAllBeers() ([]api.BeerResponse, error) {
	beers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := domain.ToDTOCollection(beers)
	return response, nil
}

func (s DefaultBeerService) GetBeer(id int) (*api.BeerResponse, error) {
	beer, err := s.repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	response := beer.ToDTO()
	return &response, nil
}

func NewBeerService(repository domain.BeerRepository) DefaultBeerService {
	return DefaultBeerService{repo: repository}
}
