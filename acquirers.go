package softpos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type AcquirerService service

type AcquirerList []Country

type Acquirer struct {
	//TODO: define structure fields
}

func (c *AcquirerService) GetList(ctx context.Context, v interface{}) (err error) {
	path := "acquirers"
	url := url.URL{Path: path}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *AcquirerService) GetDetails(ctx context.Context, code int, v interface{}) (err error) {
	path := "acquirers"
	url := url.URL{Path: fmt.Sprintf("%s/%d", path, code)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}
