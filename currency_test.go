package softpos

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

var (
	currencyRe = regexp.MustCompile(`^\/currencies\/(\d+)`)
)

func TestCurrenciesGetListMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mux.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"name":"BYN","code":933,"decimalPlaces":2,"sign":"Br"},{"name":"KZT","code":398,"decimalPlaces":2,"sign":"₸"}]`)
	})

	cl := CurrencyList{}
	err := c.CurrencyService.GetList(context.Background(), &cl)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	want := 2

	if len(cl) != want {
		t.Errorf("Countries count = %v, want %v", len(cl), want)
	}

}

func TestCurrenciesGetDetailsMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	want := 248
	mux.HandleFunc("/currencies/248", func(w http.ResponseWriter, r *http.Request) {
		if !currencyRe.MatchString(r.URL.Path) {
			t.Errorf("Bad URL got %v", r.URL.Path)
		}
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"name":"Aland Islands","nameNative":"bla bla bla","alpha2":"AX","alpha3":"ALA","code":%d}`, want)
	})

	cntr := Currency{}
	err := c.CurrencyService.GetDetails(context.Background(), want, &cntr)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if cntr.Code != want {
		t.Errorf("want %d got = %v", want, cntr.Code)
	}

}

func BenchmarkCurrencyGetListMock(b *testing.B) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
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
