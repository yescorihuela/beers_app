package domain

type BeerRepositoryStub struct {
	beers []Beer
}

func (s BeerRepositoryStub) FindAll() ([]Beer, error) {
	return s.beers, nil
}

func NewBeerRepositoryStub() BeerRepositoryStub {
	beers := []Beer{}
	return BeerRepositoryStub{beers}
}
