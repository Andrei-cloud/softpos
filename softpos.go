package softpos

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sync"
)

const (
	libraryVersion = "0.1.0"
	defaultBase    = "https://10.1.15.197:51000/api/"
	devBase        = "https://10.1.15.197:51000/api/"
	defaultUA      = "go-softpos-client/" + libraryVersion
)

func NewClient(defaultBaseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if defaultBaseURL == "" {
		defaultBaseURL = defaultBase
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		BaseURL:   baseURL,
		UserAgent: defaultUA,
		client:    httpClient,
	}

	c.CountryService = &CountryService{client: c}
	c.CurrencyService = &CurrencyService{client: c}
	c.MerchantService = &MerchantService{client: c}
	c.TerminalService = &TerminalService{client: c}
	return c
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	BaseURL   *url.URL
	UserAgent string
	apiKey    string

	CountryService  *CountryService
	CurrencyService *CurrencyService
	MerchantService *MerchantService
	TerminalService *TerminalService
}

type service struct {
	client *Client
}

func (c *Client) SetAPIKey(apiKey string) {
	c.apiKey = apiKey
}

func (c *Client) SetTransport(roundTripper http.RoundTripper) {
	c.client.Transport = roundTripper
}

// Client returns the http.Client used by this softpos client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

func (c *Client) NewRequest(method string, path url.URL, body interface{}) (*http.Request, error) {
	return c.newRequestCtx(context.Background(), method, path, body)
}

func (c *Client) newRequestCtx(ctx context.Context, method string, path url.URL, body interface{}) (*http.Request, error) {
	u := c.BaseURL.ResolveReference(&path)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Authorization", c.apiKey)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) processRequest(ctx context.Context, method string, path url.URL, body interface{}, result interface{}) error {
	req, err := c.newRequestCtx(ctx, method, path, body)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		resp := result
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			return err
		}
	case http.StatusBadRequest:
		err = ErrIncorrect
	case http.StatusUnauthorized:
		err = ErrIvalidToken
	case http.StatusForbidden:
		err = ErrNoPermission
	case http.StatusConflict:
		err = ErrConflict
	case http.StatusNotFound:
		err = ErrEntityNotFound
	default:
		err = ErrUnknown
	}

	return err
}

// func (c *Client) processBulkRequest(ctx context.Context, method string, path url.URL, params map[string]string, paramfiles map[string]string, u, f interface{}) error {
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)

// 	for p, filePath := range paramfiles {
// 		file, err := os.Open(filePath)
// 		if err != nil {
// 			return err
// 		}
// 		defer file.Close()
// 		part, err := writer.CreateFormFile(p, filepath.Base(file.Name()))
// 		if err != nil {
// 			return err
// 		}
// 		io.Copy(part, file)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for key, val := range params {
// 		_ = writer.WriteField(key, val)
// 	}

// 	writer.Close()
// 	req, err := c.newMultiPartRequestCtx(ctx, method, path, body)
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	res, err := c.client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode == http.StatusNotFound {
// 		return ErrNotFound
// 	}

// 	resp := BulkResponse{
// 		Failed:  f,
// 		Updated: u,
// 	}

// 	err = json.NewDecoder(res.Body).Decode(&resp)
// 	if err != nil {
// 		return err
// 	}

// 	if res.StatusCode == http.StatusUnauthorized {
// 		return ErrUnauthorized
// 	}
// 	if !resp.Success {
// 		err = fmt.Errorf("api err: %s", resp.Message)
// 	}
// 	return err
// }

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	return resp, err
}

// func addOptions(s string, opt interface{}) (*url.URL, error) {
// 	v := reflect.ValueOf(opt)
// 	if v.Kind() == reflect.Ptr && v.IsNil() {
// 		return &url.URL{Path: s}, nil
// 	}
// 	u, err := url.Parse(s)
// 	if err != nil {
// 		return &url.URL{Path: s}, err
// 	}
// 	vs, err := query.Values(opt)
// 	if err != nil {
// 		return &url.URL{Path: s}, err
// 	}
// 	u.RawQuery = vs.Encode()
// 	return u, nil
// }
