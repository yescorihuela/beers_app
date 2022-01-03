package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yescorihuela/beers_app/errs"
	mocks "github.com/yescorihuela/beers_app/mocks/domain"
)

func TestConvertPriceService(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockCurrencyRepo := mocks.NewMockCurrencyRepository(controller)
	serviceCurrency := NewCurrencyService(mockCurrencyRepo)

	fromCurrencyDummy := "CLP"
	toCurrencyDummy := "EUR"
	valueDummy := float32(1200.00)

	expectedValue := float32(1.20)
	// errs.NewNotFoundError("Empty model")
	mockCurrencyRepo.EXPECT().ConvertPrice(fromCurrencyDummy, toCurrencyDummy, valueDummy).Return(&expectedValue, nil)
	priceConverted, _ := serviceCurrency.ConvertPrice(fromCurrencyDummy, toCurrencyDummy, valueDummy)

	if *priceConverted != expectedValue {
		t.Error("Test failed when trying to convert currency")
	}
}

func TestConvertPriceUnavailableService(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockCurrencyRepo := mocks.NewMockCurrencyRepository(controller)
	serviceCurrency := NewCurrencyService(mockCurrencyRepo)

	fromCurrencyDummy := "CLP"
	toCurrencyDummy := "EUR"
	valueDummy := float32(1200.00)

	mockCurrencyRepo.EXPECT().ConvertPrice(fromCurrencyDummy, toCurrencyDummy, valueDummy).Return(nil, errs.NewUnexpectedError("Service Error"))
	_, err := serviceCurrency.ConvertPrice(fromCurrencyDummy, toCurrencyDummy, valueDummy)

	if err == nil {
		t.Error("Test failed when trying failure of service")
	}
}
