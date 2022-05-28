package softpos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type TerminalService service

func (c *TerminalService) GetListByMerchnat(ctx context.Context, mid string, v interface{}) (err error) {
	path := "merchants/%s/terminals"
	url := url.URL{Path: fmt.Sprintf(path, mid)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *TerminalService) GetDetailsByMerchant(ctx context.Context, mid, tid string, v interface{}) (err error) {
	path := "merchants/%s/terminals/%s"
	url := url.URL{Path: fmt.Sprintf(path, mid, tid)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *TerminalService) Create(ctx context.Context, mid string, data *Terminal, v interface{}) (err error) {
	path := "merchants/%s/terminals"
	rel := &url.URL{Path: fmt.Sprintf(path, mid)}
	if data == nil {
		return errors.New("can't create terminal on nil data")
	}

	req, err := c.client.newRequestCtx(ctx, http.MethodPost, *rel, data)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		err = json.NewDecoder(res.Body).Decode(v)
		if err != nil {
			return err
		}
	} else {
		switch res.StatusCode {
		case http.StatusBadRequest:
			err = fmt.Errorf("create terminal: %w", ErrIncorrect)
		case http.StatusUnauthorized:
			err = fmt.Errorf("create terminal: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("create terminal: %w", ErrNoPermission)
		case http.StatusConflict:
			err = fmt.Errorf("create terminal: %w", ErrConflict)
		case http.StatusNotFound:
			err = fmt.Errorf("create terminal: %w", ErrEntityNotFound)
		default:
			err = fmt.Errorf("create terminal: %w", ErrUnknown)
		}
	}

	return err
}

func (c *TerminalService) Update(ctx context.Context, ref string, data interface{}) (err error) {
	path := "terminals/%s"
	rel := &url.URL{Path: fmt.Sprintf(path, ref)}
	if data == nil {
		return errors.New("can't update on nil data")
	}

	req, err := c.client.newRequestCtx(ctx, http.MethodPatch, *rel, data)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		err = nil
	} else {
		switch res.StatusCode {
		case http.StatusBadRequest:
			err = fmt.Errorf("status terminal: %w", ErrIncorrect)
		case http.StatusUnauthorized:
			err = fmt.Errorf("status terminal: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("status terminal: %w", ErrNoPermission)
		case http.StatusNotFound:
			err = fmt.Errorf("status terminal: %w", ErrEntityNotFound)
		case http.StatusConflict:
			err = fmt.Errorf("create merchant: %w", ErrConflict)
		default:
			err = fmt.Errorf("status terminal: %w", ErrUnknown)
		}
	}

	return err
}

func (c *TerminalService) ChangeStatus(ctx context.Context, mid, tid string, state interface{}) (err error) {
	path := "merchants/%s/terminals/%s/status"
	rel := &url.URL{Path: fmt.Sprintf(path, mid, tid)}
	if state == nil {
		return errors.New("can't cahnge status on nil data")
	}

	req, err := c.client.newRequestCtx(ctx, http.MethodPut, *rel, state)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		err = nil
	} else {
		switch res.StatusCode {
		case http.StatusBadRequest:
			err = fmt.Errorf("status terminal: %w", ErrIncorrect)
		case http.StatusUnauthorized:
			err = fmt.Errorf("status terminal: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("status terminal: %w", ErrNoPermission)
		case http.StatusNotFound:
			err = fmt.Errorf("status terminal: %w", ErrEntityNotFound)
		default:
			err = fmt.Errorf("status terminal: %w", ErrUnknown)
		}
	}

	return err
}
