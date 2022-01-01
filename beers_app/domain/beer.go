package domain

import (
	"time"

	"github.com/yescorihuela/beers_app/api"
	"github.com/yescorihuela/beers_app/errs"
)

type Beer struct {
	Id        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name"`
	Brewery   string     `json:"brewery"`
	Country   string     `json:"country"`
	Price     float32    `json:"price" gorm:"precision:16;scale:2"`
	Currency  string     `json:"currency"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"autoUpdateTime:milli"`
}

func NewBeer(name, brewery, country, currency string, price float32) Beer {
	return Beer{
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}
}

func (b Beer) ToDTO() api.BeerResponse {
	return api.BeerResponse{
		Id:       b.Id,
		Name:     b.Name,
		Brewery:  b.Brewery,
		Country:  b.Country,
		Price:    b.Price,
		Currency: b.Currency,
	}
}

func (b Beer) ToTotalPriceDTO(quantity float32) api.BeerBoxTotalPrice {
	return api.BeerBoxTotalPrice{
		Price: b.Price * quantity,
	}
}

func ToDTOCollection(b []Beer) []api.BeerResponse {
	var beers []api.BeerResponse
	for _, beer := range b {
		beers = append(beers, beer.ToDTO())
	}
	return beers
}

type BeerRepository interface {
	FindAll() ([]Beer, *errs.AppError)
	FindOne(int) (*Beer, *errs.AppError)
	Create(beer Beer) (*Beer, *errs.AppError)
}
