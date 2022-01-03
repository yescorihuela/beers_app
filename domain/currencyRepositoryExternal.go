package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/yescorihuela/beers_app/errs"
)

const bridgeCurrencyPrefix = "USD"

type CurrencyRepositoryExternal struct {
	client *http.Client
}

func NewCurrencyRepositoryExternal(client *http.Client) CurrencyRepositoryExternal {
	return CurrencyRepositoryExternal{client: client}
}

func (cre CurrencyRepositoryExternal) ConvertPrice(fromCurrency string, toCurrency string, value float32) (*float32, *errs.AppError) {

	currency, err := executeRequest(cre, fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}

	if currency.ValueFrom == 0 {
		return nil, &errs.AppError{Message: "Division by zero in ValueFrom key", Code: http.StatusServiceUnavailable}
	}

	if currency.ValueTo == 0 {
		return nil, &errs.AppError{Message: "Division by zero in ValueFrom key", Code: http.StatusServiceUnavailable}
	}
	convertedValue := (value / currency.ValueFrom) * currency.ValueTo
	return &convertedValue, nil
}

func executeRequest(cre CurrencyRepositoryExternal, fromCurrency, toCurrency string) (*Currency, *errs.AppError) {
	var currency Currency
	fromCurrencyAbbr := fmt.Sprintf("%s%s", bridgeCurrencyPrefix, fromCurrency)
	toCurrencyAbbr := fmt.Sprintf("%s%s", bridgeCurrencyPrefix, toCurrency)
	baseUrl := os.Getenv("CURRENCY_LAYER_API_HOST")
	accessKey := os.Getenv("CURRENCY_LAYER_API_KEY")
	url := fmt.Sprintf("%s/live?access_key=%s&currencies=%s,%s&format=1", baseUrl, accessKey, fromCurrency, toCurrency)
	response, err := cre.client.Get(url)
	if err != nil {

		return nil, errs.NewUnexpectedError(err.Error())
	}
	defer response.Body.Close()

	conversion := struct {
		Quotes map[string]float32 `json:"quotes"`
	}{}
	json.NewDecoder(response.Body).Decode(&conversion)

	currency.setValues(conversion.Quotes, fromCurrencyAbbr, toCurrencyAbbr)

	return &currency, nil
}

func (c *Currency) setValues(data map[string]float32, fromKey, toKey string) {
	c.ValueFrom = data[fromKey]
	c.ValueTo = data[toKey]
}
