package dtos

type PaymentMethodRequestDTO struct {
	CODE         string `json:"code"`
	JOURNAL      string `json:"journal"`
	CURRENCYCODE string `json:"currency_code"`
	DESCRIPTION  string `json:"description"`
}
