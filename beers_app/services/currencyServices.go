package services

import (
	"github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/errs"
)

type CurrencyService interface {
	ConvertPrice(fromCurrency string, toCurrency string, value float32) (*float32, *errs.AppError)
}

type DefaultCurrencyService struct {
	repo domain.CurrencyRepository
}

func NewCurrencyService(ds domain.CurrencyRepository) DefaultCurrencyService {
	return DefaultCurrencyService{repo: ds}
}

func (d DefaultCurrencyService) ConvertPrice(fromCurrency string, toCurrency string, value float32) (*float32, *errs.AppError) {
	convertedPrice, err := d.repo.ConvertPrice(fromCurrency, toCurrency, value)
	if err != nil {
		return nil, err
	}
	return convertedPrice, nil
}
