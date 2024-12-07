package epayment

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type CryptogramResponse struct {
	Hpan       string `json:"hpan"`
	ExpDate    string `json:"expDate"`
	Cvc        string `json:"cvc"`
	TerminalId string `json:"terminalId"`
}

type EpaymentResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	InvoiceID string  `json:"invoice_id"`
}
