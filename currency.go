package softpos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type CurrencyService service

type CurrencyList []Currency

type Currency struct {
	Name          string `json:"name"`
	Code          int    `json:"code"`
	DecimalPlaces int    `json:"decimalPlaces"`
	Sign          string `json:"sign"`
}

func (c *CurrencyService) GetList(ctx context.Context, v interface{}) (err error) {
	path := "currencies"
	url := url.URL{Path: path}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *CurrencyService) GetDetails(ctx context.Context, code int, v interface{}) (err error) {
	path := "currencies"
	url := url.URL{Path: fmt.Sprintf("%s/%d", path, code)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}
