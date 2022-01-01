package api

import "github.com/yescorihuela/beers_app/errs"

type NewBeerRequest struct {
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}

func (nbr NewBeerRequest) Validate() *errs.AppError {
	if nbr.Name == "" {
		return errs.NewValidationError("Name is required")
	}

	if nbr.Brewery == "" {
		return errs.NewValidationError("Brewery is required")
	}

	if nbr.Country == "" {
		return errs.NewValidationError("Country is required")
	}

	if nbr.Currency == "" {
		return errs.NewValidationError("Currency is required")
	}

	return nil
}
