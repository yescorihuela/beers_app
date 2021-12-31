package domain

import (
	"time"

	"github.com/yescorihuela/beers_app/api"
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

func ToDTOCollection(b []Beer) []api.BeerResponse {
	beers := make([]api.BeerResponse, len(b))
	for _, beer := range b {
		beers = append(beers, beer.ToDTO())
	}
	return beers
}

type BeerRepository interface {
	FindAll() ([]Beer, error)
	// These methods will be used in the real implementation
	// FindOne(beer_id int) (Beer, error)
	// Create(beer Beer) error
}
