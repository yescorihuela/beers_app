package services

import (
	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/errs"
)

type BeerService interface {
	GetAllBeers() ([]api.BeerResponse, *errs.AppError)
	GetBeer(id int) (*api.BeerResponse, *errs.AppError)
	Create(req api.NewBeerRequest) (*api.BeerResponse, *errs.AppError)
	GetBeerByBox(c CurrencyService, id int, quantity float32, toCurrency string) (*api.BeerBoxTotalPrice, *errs.AppError)
}

type DefaultBeerService struct {
	repo domain.BeerRepository
}

func (s DefaultBeerService) GetAllBeers() ([]api.BeerResponse, *errs.AppError) {
	beers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := domain.ToDTOCollection(beers)
	return response, nil
}

func (s DefaultBeerService) GetBeer(id int) (*api.BeerResponse, *errs.AppError) {
	beer, err := s.repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	response := beer.ToDTO()
	return &response, nil
}

func (s DefaultBeerService) Create(req api.NewBeerRequest) (*api.BeerResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	beer := domain.NewBeer(req.Name, req.Brewery, req.Country, req.Currency, req.Price)
	if newBeer, err := s.repo.Create(beer); err != nil {
		return nil, err
	} else {
		response := newBeer.ToDTO()
		return &response, nil
	}
}

func (s DefaultBeerService) GetBeerByBox(c CurrencyService, id int, quantity float32, toCurrency string) (*api.BeerBoxTotalPrice, *errs.AppError) {
	beer, err := s.repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	convertedPrice, err := c.ConvertPrice(beer.Currency, toCurrency, beer.Price)
	if err != nil {
		return nil, err
	}
	response := beer.ToTotalPriceDTO(quantity, *convertedPrice)
	return &response, nil
}

func NewBeerService(repository domain.BeerRepository) DefaultBeerService {
	return DefaultBeerService{repo: repository}
}
