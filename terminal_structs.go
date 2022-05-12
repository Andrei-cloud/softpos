package softpos

import "time"

type Terminal struct {
	TerminalID  string `json:"terminalId"`
	MerchantRef string `json:"merchantRef,omitempty"`
	Currency    int    `json:"currency"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Profile     string `json:"profile"`
	Name        string `json:"name"`
	Mcc         int    `json:"mcc"`
	State       string `json:"state"`
	Note        string `json:"note,omitempty"`
	Language    string `json:"language"`
}

type TemrinalDetails struct {
	CurrencyName         string `json:"currencyName"`
	TerminalCurrencyName string `json:"terminalCurrencyName"`
	Currency             int    `json:"currency"`
	Country              int    `json:"country"`
	Mcc                  int    `json:"mcc"`
	TerminalMcc          string `json:"terminalMcc"`
	Profile              string `json:"profile"`
	Language             string `json:"language"`
	Merchant             struct {
		CurrencyName       string    `json:"currencyName"`
		AcquirerName       string    `json:"acquirerName"`
		CountryName        string    `json:"countryName"`
		CountryNativeName  string    `json:"countryNativeName"`
		Mcc                string    `json:"mcc"`
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
		Language           string    `json:"language"`
		Profile            string    `json:"profile"`
		Flags              string    `json:"flags"`
	} `json:"merchant"`
	Preferences             []Preferences `json:"preferences"`
	InputMethods            []string      `json:"inputMethods"`
	State                   string        `json:"state"`
	Reference               string        `json:"reference"`
	TerminalID              string        `json:"terminalId"`
	CurrentBatchRef         string        `json:"currentBatchRef"`
	Keys                    []Keys        `json:"keys"`
	Created                 time.Time     `json:"created"`
	Updated                 time.Time     `json:"updated"`
	MasterKeyID             string        `json:"masterKeyId"`
	KeysConfirmed           bool          `json:"keysConfirmed"`
	OperationSequenceNumber int           `json:"operationSequenceNumber"`
	Phone                   string        `json:"phone"`
	TerminalProfile         string        `json:"terminalProfile"`
	Name                    string        `json:"name"`
	Email                   string        `json:"email"`
	TerminalCurrency        int           `json:"terminalCurrency"`
	SequenceNumber          int           `json:"sequenceNumber"`
	TerminalLanguage        string        `json:"terminalLanguage"`
}

type Preferences struct {
	Tag           string `json:"tag"`
	Value         bool   `json:"value"`
	Description   string `json:"description"`
	PaymentSystem string `json:"paymentSystem"`
	Type          string `json:"type"`
}
type Keys struct {
	KeyType       string `json:"keyType"`
	Encoding      string `json:"encoding"`
	KeyValue      string `json:"keyValue"`
	KeyCheckValue string `json:"keyCheckValue"`
	KeyID         int    `json:"keyId"`
}
