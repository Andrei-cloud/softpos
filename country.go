package softpos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type CountryService service

type CountryList []Country

type Country struct {
	Name       string `json:"name"`
	NameNative string `json:"nameNative"`
	Alpha2     string `json:"alpha2"`
	Alpha3     string `json:"alpha3"`
	Code       int    `json:"code"`
}

func (c *CountryService) GetList(ctx context.Context, v interface{}) (err error) {
	path := "countries"
	url := url.URL{Path: path}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *CountryService) GetDetails(ctx context.Context, code int, v interface{}) (err error) {
	path := "countries"
	url := url.URL{Path: fmt.Sprintf("%s/%d", path, code)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}
