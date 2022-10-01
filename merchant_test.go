package softpos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var (
	merchantRe = regexp.MustCompile(`^\/merchants\/([a-zA-Z0-9]+)`)
)

func TestMerchantGetListMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mux.HandleFunc("/merchants", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"index":0,"totalPages":1,"count":2,"totalCount":2,"perPage":10,"offset":0,"items":[{"currencyName":"QAR","acquirerName":"CBQ","countryName":"Qatar","countryNativeName":"قطر","mcc":"5910","state":"Active","reference":"32c24e1d-3346-4c5e-a15a-54ffe4e54712","merchantId":"600086900","isLocationRequired":false,"name":"MERCHANT UAT","taxRefNumber":"6453746","country":634,"city":"DOHA","region":"DOHA","address":"wastbay ","postalCode":"50000","phone":"+97466667777","email":"zakaria.taqui@cbq.qa","created":"2022-02-20T11:58:14.483339Z","updated":"2022-02-20T11:58:14.483339Z","acquirer":"cbq","currency":634,"language":"en","profile":"default","flags":"None"},{"currencyName":"QAR","acquirerName":"CBQ","countryName":"Qatar","countryNativeName":"قطر","mcc":"5910","state":"Active","reference":"643efb2d-adfc-4674-aacb-2faa4667e97d","merchantId":"999700163","isLocationRequired":false,"name":"merchnat","taxRefNumber":"XYZ","country":634,"city":"DOHA","region":"MIDDLE EAST","address":"address","postalCode":"3232","phone":"44448888","email":"user1@example.com","created":"2022-02-28T07:25:42.076142Z","updated":"2022-02-28T07:25:42.076142Z","acquirer":"cbq","currency":634,"language":"en","profile":"default","flags":"None"}]}`)
	})

	ml := MerchnatList{}
	err := c.MerchantService.GetList(context.Background(), &ml)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	want := 2

	if ml.Count != want {
		t.Errorf("Merchants count = %v, want %v", ml.Count, want)
	}

}

func TestMerchnatGetDetailsMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	want := "600068900"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s", want), func(w http.ResponseWriter, r *http.Request) {
		if !merchantRe.MatchString(r.URL.Path) {
			t.Errorf("Bad URL got %v", r.URL.Path)
		}
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"state":"Active","reference":"32c24e1d-3346-4c5e-a15a-54ffe4e54712","merchantId":"%s","isLocationRequired":false,"name":"MERCHANT UAT","taxRefNumber":"6453746","country":634,"city":"DOHA","region":"DOHA","address":"wastbay ","postalCode":"50000","phone":"+97466667777","email":"merchnat@example.com","created":"2022-02-20T11:58:14.483339Z","updated":"2022-02-20T11:58:14.483339Z","acquirer":"cbq","currency":634,"mcc":5910,"language":"en","profile":"default","flags":"None"}`, want)
	})

	merchnat := MerchantDetails{}
	err := c.MerchantService.GetDetails(context.Background(), want, &merchnat)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if merchnat.MerchantID != want {
		t.Errorf("want %v got = %v", want, merchnat.MerchantID)
	}

}

func TestMerchnatCreateMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	merchnat := MerchantDetails{
		State:              "Active",
		MerchantID:         "750074750",
		IsLocationRequired: false,
		Name:               "Merchant Test",
		TaxRefNumber:       "X505",
		Country:            634,
		City:               "Doha",
		Region:             "Middle east",
		Address:            "west bay, doha",
		PostalCode:         "12300",
		Phone:              "+97465743782",
		Email:              "zak.exemple@cbq.qa",
		Acquirer:           "cbq",
		Currency:           634,
		Mcc:                5812,
		Language:           "en",
		Profile:            "default",
	}

	want := "3ddf6776-b872-4053-8391-a6c3db5fb008"
	mux.HandleFunc("/merchants", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		m := MerchantDetails{}
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			t.Errorf("Error occured = %v", err)
		}
		if !reflect.DeepEqual(merchnat, m) {
			t.Error("Error occured: entities are NOT equal")
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"reference":"%s"}`, want)
	})

	ref := CreateResponse{}
	err := c.MerchantService.Create(context.Background(), &merchnat, &ref)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if ref.Reference != want {
		t.Errorf("want %v got = %v", want, ref.Reference)
	}
}

func TestMerchnatCreateConflictMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	merchnat := MerchantDetails{
		State:              "Active",
		MerchantID:         "750074750",
		IsLocationRequired: false,
		Name:               "Merchant Test",
		TaxRefNumber:       "X505",
		Country:            634,
		City:               "Doha",
		Region:             "Middle east",
		Address:            "west bay, doha",
		PostalCode:         "12300",
		Phone:              "+97465743782",
		Email:              "zak.exemple@cbq.qa",
		Acquirer:           "cbq",
		Currency:           634,
		Mcc:                5812,
		Language:           "en",
		Profile:            "default",
	}

	want := "3ddf6776-b872-4053-8391-a6c3db5fb008"
	mux.HandleFunc("/merchants", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, `{"reason":"%s", "field":"MID", "value": "%s" }`, want, merchnat.MerchantID)
	})

	ref := CreateResponse{}
	err := c.MerchantService.Create(context.Background(), &merchnat, &ref)
	if !strings.Contains(err.Error(), want) {
		t.Errorf("Error occured = %v", err)
	}
}

func TestMerchnatChangeStatusMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	state := struct {
		State string `json:"state"`
		Note  string `json:"note,omitempty"`
	}{"Active", "activate terminal"}

	want := "750074750"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s/status", want), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		s := make(map[string]string)
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			t.Errorf("Error occured = %v", err)
		}
		if s["state"] != state.State {
			t.Errorf("Error occured = want %v, got %v", state.State, s["state"])
		}
		w.WriteHeader(http.StatusOK)
	})

	err := c.MerchantService.ChangeStatus(context.Background(), want, &state)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}
}

func BenchmarkMerchnatGetListMock(b *testing.B) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/merchants", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"BYN","code":933,"decimalPlaces":2,"sign":"Br"},{"name":"KZT","code":398,"decimalPlaces":2,"sign":"₸"}]`)
	})

	cl := CurrencyList{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		err := c.CurrencyService.GetList(context.Background(), &cl)
		if err != nil {
			b.Errorf("Error occured = %v", err)
		}
	}
}
