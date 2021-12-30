package domain

type Beer struct {
	Id   string
	Name string
}

type BeerRepository interface {
	FindAll() ([]Beer, error)
	// These methods will be used in the real implementation
	// FindOne(beer_id int) (Beer, error)
	// Create(beer Beer) error
}
