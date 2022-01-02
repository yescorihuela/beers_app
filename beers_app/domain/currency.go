package domain

import "github.com/yescorihuela/beers_app/errs"

type Currency struct {
	Source    string  `json:"source,omitempty"`
	ValueFrom float32 `json:"value_from"`
	ValueTo   float32 `json:"value_to"`
}

type CurrencyRepository interface {
	ConvertPrice(fromCurrency string, toCurrency string, value float32) (*float32, *errs.AppError)
}
