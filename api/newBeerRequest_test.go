package api

import (
	"net/http"
	"testing"
)

func TestNewBeerRequest(t *testing.T) {
	request := NewBeerRequest{Id: 0}

	err := request.Validate()
	if err.Message != "Id is required" {
		t.Error("Invalid message while testing beer ID")
	}
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer ID")
	}

	// Set beer ID after test beer ID
	request.Id = 1

	err = request.Validate()
	if err.Message != "Name is required" {
		t.Error("Invalid message while testing beer name")
	}
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer name")
	}

	// Set beer name after test beer name
	request.Name = "Torobayo"

	err = request.Validate()
	if err.Message != "Brewery is required" {
		t.Error("Invalid message while testing beer Brewery")
	}
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer Brewery")
	}

	// Set beer brewery after test beer brewery
	request.Brewery = "Kunstmann"

	err = request.Validate()
	if err.Message != "Country is required" {
		t.Error("Invalid message while testing beer Country")
	}
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer Country")
	}

	// Set beer country after test beer country
	request.Country = "Chile"

	err = request.Validate()
	if err.Message != "Price must be higher than zero" {
		t.Error("Invalid message while testing beer Price")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer Price")
	}

	// Set beer price after test beer price
	request.Price = 1200.00

	err = request.Validate()
	if err.Message != "Currency is required" {
		t.Error("Invalid message while testing beer Price")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid HTTP code while testing beer Price")
	}

	// Set beer currency after test beer currency
	request.Currency = "CLP"

	err = request.Validate()
	if err != nil {
		t.Error("Invalid beer request")
	}

}
