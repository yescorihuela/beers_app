package services

import "github.com/yescorihuela/beers_app/domain"

type BeerService interface {
	GetAllBeers() ([]domain.Beer, error)
}

type DefaultBeerService struct {
	repo domain.BeerRepository
}

func (s DefaultBeerService) GetAllBeers() ([]domain.Beer, error) {
	return s.repo.FindAll()
}

func NewBeerService(repository domain.BeerRepository) DefaultBeerService {
	return DefaultBeerService{repo: repository}
}
