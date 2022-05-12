package softpos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"testing"
)

var (
	TerminalRe = regexp.MustCompile(`^\/merchants\/([a-zA-Z0-9]+)\/terminals`)
)

func TestTerminalGetListByMerchantMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mid := "600086900"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s/terminals", mid), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"currencyName":"QAR","terminalCurrencyName":"QAR","currency":634,"country":634,"mcc":5910,"terminalMcc":"5910","profile":"default","language":"en","merchant":{"currencyName":"QAR","acquirerName":"CBQ","countryName":"Qatar","countryNativeName":"قطر","mcc":"5910","state":"Active","reference":"32c24e1d-3346-4c5e-a15a-54ffe4e54712","merchantId":"600086900","isLocationRequired":false,"name":"MERCHANT UAT","taxRefNumber":"6453746","country":634,"city":"DOHA","region":"DOHA","address":"wastbay ","postalCode":"50000","phone":"+974661642269","email":"zakaria.taqui@cbq.qa","created":"2022-02-20T11:58:14.483339Z","updated":"2022-02-20T11:58:14.483339Z","acquirer":"cbq","currency":634,"language":"en","profile":"default","flags":"None"},"preferences":[],"inputMethods":[],"state":"Active","reference":"c5600602-a2b1-48f2-a1d7-9475c4454191","terminalId":"66770057","currentBatchRef":"50e8922d-c9a8-436e-9b5b-d511fedf8a8f","keys":[{"keyType":"TPK","encoding":"LMK","keyValue":"UAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","keyCheckValue":"083CA1","keyId":0}],"created":"2022-03-01T08:51:57.229638Z","updated":"2022-03-07T14:37:43.839132Z","masterKeyId":"0","keysConfirmed":true,"operationSequenceNumber":0,"phone":"+97466164269","terminalProfile":"default","name":"zaka","email":"santosh.babar@cbq.qa","terminalCurrency":634,"sequenceNumber":0,"terminalLanguage":"en"},{"currencyName":"QAR","terminalCurrencyName":"QAR","currency":634,"country":634,"mcc":5810,"terminalMcc":"5810","profile":"default","language":"en","merchant":{"currencyName":"QAR","acquirerName":"CBQ","countryName":"Qatar","countryNativeName":"قطر","mcc":"5910","state":"Active","reference":"32c24e1d-3346-4c5e-a15a-54ffe4e54712","merchantId":"600086900","isLocationRequired":false,"name":"MERCHANT UAT","taxRefNumber":"6453746","country":634,"city":"DOHA","region":"DOHA","address":"wastbay ","postalCode":"50000","phone":"+974661642269","email":"zakaria.taqui@cbq.qa","created":"2022-02-20T11:58:14.483339Z","updated":"2022-02-20T11:58:14.483339Z","acquirer":"cbq","currency":634,"language":"en","profile":"default","flags":"None"},"preferences":[],"inputMethods":[],"state":"Active","reference":"d1033a63-3ed2-4b40-813f-4bb515966c7a","terminalId":"66770056","currentBatchRef":"3e05c07d-3e46-4262-bc72-3ae15a2066fe","keys":[{"keyType":"TMK","encoding":"LMK","keyValue":"UAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","keyCheckValue":"BAA397","keyId":0}],"created":"2022-03-01T08:50:12.986945Z","updated":"2022-03-14T12:30:54.694094Z","masterKeyId":"0","keysConfirmed":true,"operationSequenceNumber":0,"phone":"+97466164269","terminalProfile":"default","name":"zaka","email":"arun@appknox.com","terminalCurrency":634,"sequenceNumber":1,"terminalLanguage":"en"}]`)
	})

	tl := []TemrinalDetails{}
	err := c.TerminalService.GetListByMerchnat(context.Background(), mid, &tl)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	want := 2

	if len(tl) != want {
		t.Errorf("Tetminals count = %v, want %v", len(tl), want)
	}

}

func TestTerminalGetDetailsMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	mid := "600086900"
	want := "66770050"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s/terminals/%s", mid, want), func(w http.ResponseWriter, r *http.Request) {
		if !merchantRe.MatchString(r.URL.Path) {
			t.Errorf("Bad URL got %v", r.URL.Path)
		}
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"currencyName":"QAR","terminalCurrencyName":"QAR","currency":634,"country":634,"mcc":5812,"terminalMcc":"5812","profile":"default","language":"en","merchant":{"currencyName":"QAR","acquirerName":"CBQ","countryName":"Qatar","countryNativeName":"قطر","mcc":"5910","state":"Active","reference":"32c24e1d-3346-4c5e-a15a-54ffe4e54712","merchantId":"600086900","isLocationRequired":false,"name":"MERCHANT UAT","taxRefNumber":"6453746","country":634,"city":"DOHA","region":"DOHA","address":"wastbay ","postalCode":"50000","phone":"+974661642269","email":"zakaria.taqui@cbq.qa","created":"2022-02-20T11:58:14.483339Z","updated":"2022-02-20T11:58:14.483339Z","acquirer":"cbq","currency":634,"language":"en","profile":"default","flags":"None"},"preferences":[{"tag":"readerCvmRequiredLimitEnabled","value":true,"description":"CVM Required Limit Enabled","paymentSystem":"VISA","type":"Boolean"}],"inputMethods":["Contactless"],"state":"Active","reference":"b59b8acc-21ae-45c8-84d8-ce044c9bfec3","terminalId":"%s","currentBatchRef":"56e336d0-3fa6-4597-9b1e-590314e4b196","keys":[{"keyType":"TMK","encoding":"LMK","keyValue":"UAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","keyCheckValue":"BAA397","keyId":0}],"created":"2022-03-28T09:16:29.293528Z","updated":"2022-03-28T09:16:29.307507Z","masterKeyId":"0","keysConfirmed":true,"operationSequenceNumber":0,"phone":"+97453627564","terminalProfile":"default","name":"zak termname","email":"zak.example@cbq.qa","terminalCurrency":634,"sequenceNumber":0,"terminalLanguage":"en"}`, want)
	})

	term := TemrinalDetails{}
	err := c.TerminalService.GetDetailsByMerchant(context.Background(), mid, want, &term)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if term.TerminalID != want {
		t.Errorf("want %v got = %v", want, term.TerminalID)
	}

}

func TestTerminalCreateMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	terminal := Terminal{
		TerminalID: "66770050",
		Currency:   634,
		Phone:      "+97453627564",
		Email:      "zak.example@cbq.qa",
		Profile:    "default",
		Name:       "zak termname",
		Mcc:        5812,
		State:      "Active",
		Note:       "",
		Language:   "en",
	}

	mid := "600086900"
	want := "b59b8acc-21ae-45c8-84d8-ce044c9bfec3"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s/terminals", mid), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		ter := Terminal{}
		if err := json.NewDecoder(r.Body).Decode(&ter); err != nil {
			t.Errorf("Error occured = %v", err)
		}
		if !reflect.DeepEqual(terminal, ter) {
			t.Error("Error occured: entities are NOT equal")
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"reference":"%s"}`, want)
	})

	ref := CreateResponse{}
	err := c.TerminalService.Create(context.Background(), mid, &terminal, &ref)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}

	if ref.Reference != want {
		t.Errorf("want %v got = %v", want, ref.Reference)
	}
}

func TestTerminaltChangeStatusMock(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()

	c.client.Transport = LoggingRoundTripper{http.DefaultTransport}

	state := struct {
		State string `json:"state"`
		Note  string `json:"note,omitempty"`
	}{"Active", "activate terminal"}

	mid := "600086900"
	tid := "66770050"
	mux.HandleFunc(fmt.Sprintf("/merchants/%s/terminals/%s/status", mid, tid), func(w http.ResponseWriter, r *http.Request) {
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

	err := c.TerminalService.ChangeStatus(context.Background(), mid, tid, &state)
	if err != nil {
		t.Errorf("Error occured = %v", err)
	}
}
