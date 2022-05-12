package softpos

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

var (
	countriesRe = regexp.MustCompile(`^\/countries\/(\d+)`)
)

func TestCountriesGetListMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mux.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"name":"Afghanistan","nameNative":"balblabla","alpha2":"AF","alpha3":"AFG","code":4},{"name":"Aland Islands","nameNative":"bla bla bla","alpha2":"AX","alpha3":"ALA","code":248}]`)
	})

	cl := CountryList{}
	err := c.CountryService.GetList(context.Background(), &cl)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	want := 2

	if len(cl) != want {
		t.Errorf("Countries count = %v, want %v", len(cl), want)
	}

}

func TestCountriesGetDetailsMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	want := 248
	mux.HandleFunc("/countries/248", func(w http.ResponseWriter, r *http.Request) {
		if !countriesRe.MatchString(r.URL.Path) {
			t.Errorf("Bad URL got %v", r.URL.Path)
		}
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"name":"Aland Islands","nameNative":"bla bla bla","alpha2":"AX","alpha3":"ALA","code":%d}`, want)
	})

	cntr := Country{}
	err := c.CountryService.GetDetails(context.Background(), want, &cntr)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if cntr.Code != want {
		t.Errorf("want %d got = %v", want, cntr.Code)
	}

}

func BenchmarkCountriesGetListMock(b *testing.B) {
	c, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"Afghanistan","nameNative":"balblabla","alpha2":"AF","alpha3":"AFG","code":4},{"name":"Aland Islands","nameNative":"bla bla bla","alpha2":"AX","alpha3":"ALA","code":248}]`)
	})

	cl := CountryList{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		err := c.CountryService.GetList(context.Background(), &cl)
		if err != nil {
			b.Errorf("Error occured = %v", err)
		}
	}
}
