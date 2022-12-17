package softpos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MerchantService service

type MerchnatList struct {
	Index      int `json:"index"`
	TotalPages int `json:"totalPages"`
	Count      int `json:"count"`
	TotalCount int `json:"totalCount"`
	PerPage    int `json:"perPage"`
	Offset     int `json:"offset"`
	Items      []struct {
		State              string    `json:"state"`
		Reference          string    `json:"reference"`
		MerchantID         string    `json:"merchantId"`
		IsLocationRequired bool      `json:"isLocationRequired"`
		Name               string    `json:"name"`
		TaxRefNumber       string    `json:"taxRefNumber"`
		Country            int       `json:"country"`
		City               string    `json:"city"`
		Region             string    `json:"region"`
		Address            string    `json:"address"`
		PostalCode         string    `json:"postalCode"`
		Phone              string    `json:"phone"`
		Email              string    `json:"email"`
		Created            time.Time `json:"created"`
		Updated            time.Time `json:"updated"`
		Acquirer           string    `json:"acquirer"`
		Currency           int       `json:"currency"`
		Mcc                string    `json:"mcc"`
		Language           string    `json:"language"`
		Profile            string    `json:"profile"`
		Flags              string    `json:"flags"`
	} `json:"items"`
}

type MerchantDetails struct {
	CurrencyName       string `json:"currencyName,omitempty"`
	AcquirerName       string `json:"acquirerName,omitempty"`
	CountryName        string `json:"countryName,omitempty"`
	CountryNativeName  string `json:"countryNativeName,omitempty"`
	Mcc                string `json:"mcc,omitempty"`
	State              string `json:"state,omitempty"`
	Reference          string `json:"reference,omitempty"`
	MerchantID         string `json:"merchantId,omitempty"`
	IsLocationRequired bool   `json:"isLocationRequired,omitempty"`
	Name               string `json:"name,omitempty"`
	TaxRefNumber       string `json:"taxRefNumber,omitempty"`
	Country            int    `json:"country,omitempty"`
	City               string `json:"city,omitempty"`
	Region             string `json:"region,omitempty"`
	Address            string `json:"address,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Phone              string `json:"phone,omitempty"`
	Email              string `json:"email,omitempty"`
	Created            string `json:"created,omitempty"`
	Updated            string `json:"updated,omitempty"`
	Acquirer           string `json:"acquirer,omitempty"`
	Currency           int    `json:"currency,omitempty"`
	Language           string `json:"language,omitempty"`
	Profile            string `json:"profile,omitempty"`
	Flags              string `json:"flags,omitempty"`
}

func (c *MerchantService) GetList(ctx context.Context, v interface{}) (err error) {
	path := "merchants"
	url := url.URL{Path: path}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *MerchantService) GetDetails(ctx context.Context, mid string, v interface{}) (err error) {
	path := "merchants"
	url := url.URL{Path: fmt.Sprintf("%s/%s", path, mid)}
	return c.client.processRequest(ctx, http.MethodGet, url, nil, v)
}

func (c *MerchantService) Create(ctx context.Context, data *MerchantDetails, v interface{}) (err error) {
	path := "merchants"
	rel := &url.URL{Path: path}
	if data == nil {
		return errors.New("can't create merchant on nil data")
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
			err = fmt.Errorf("create merchant: %w", ErrAcqNotExist)

		case http.StatusUnauthorized:
			err = fmt.Errorf("create merchant: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("create merchant: %w", ErrNoPermission)
		case http.StatusConflict:
			reason := &conflict{}
			err = json.NewDecoder(res.Body).Decode(reason)
			if err != nil {
				return err
			}
			apierr := fmt.Sprintf("%s:%s:%s:%d", reason.Reason, reason.Field, reason.Value, reason.Type)
			err = fmt.Errorf("create merchant: %w", fmt.Errorf(apierr+"- %w", ErrConflict))
		default:
			err = fmt.Errorf("create merchant: %w", ErrUnknown)
		}
	}

	return err
}

func (c *MerchantService) Update(ctx context.Context, ref string, data interface{}) (err error) {
	path := "merchants/%s"
	rel := &url.URL{Path: fmt.Sprintf(path, ref)}
	if data == nil {
		return errors.New("can't update status on nil data")
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
			err = fmt.Errorf("status merchant: %w", ErrAcqNotExist)
		case http.StatusUnauthorized:
			err = fmt.Errorf("status merchant: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("status merchant: %w", ErrNoPermission)
		case http.StatusNotFound:
			err = fmt.Errorf("status merchant: %w", ErrEntityNotFound)
		case http.StatusConflict:
			err = fmt.Errorf("create merchant: %w", ErrConflict)
		default:
			err = fmt.Errorf("status merchant: %w", ErrUnknown)
		}
	}

	return err
}

func (c *MerchantService) ChangeStatus(ctx context.Context, mid string, state interface{}) (err error) {
	path := "merchants/%s/status"
	rel := &url.URL{Path: fmt.Sprintf(path, mid)}
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
			err = fmt.Errorf("status merchant: %w", ErrAcqNotExist)
		case http.StatusUnauthorized:
			err = fmt.Errorf("status merchant: %w", ErrIvalidToken)
		case http.StatusForbidden:
			err = fmt.Errorf("status merchant: %w", ErrNoPermission)
		case http.StatusNotFound:
			err = fmt.Errorf("status merchant: %w", ErrEntityNotFound)
		default:
			err = fmt.Errorf("status merchant: %w", ErrUnknown)
		}
	}

	return err
}
