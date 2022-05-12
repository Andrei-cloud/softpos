package softpos

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetListMock_withLogging(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mux.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"name":"Afghanistan","nameNative":"balblabla","alpha2":"AF","alpha3":"AFG","code":4},{"name":"Aland Islands","nameNative":"bla bla bla","alpha2":"AX","alpha3":"ALA","code":248}]`)
	})

	tl := CountryList{}
	err := c.CountryService.GetList(context.Background(), &tl)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	want := 2

	if len(tl) != want {
		t.Errorf("Countries count = %v, want %v", len(tl), want)
	}

}
