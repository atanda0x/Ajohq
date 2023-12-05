package util

const (
	XOF = "XOF"
	XAF = "XAF"
	ZAR = "ZAR"
	EUR = "EUR"
	CNY = "CNY"
	JPY = "JPY"
	INR = "INR"
	GBP = "GBP"
	CHF = "CHF"
	USD = "USD"
	CAD = "CAD"
	MXN = "MXN"
	AUD = "AUD"
	NZD = "NZD"
	BRL = "BRL"
	ARS = "ARS"
	CLP = "CLP"
	NGN = "NGN"
)

// IsSupportedCurrency returns true if the currency is suppoerted
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, USD, NGN:
		return true
	}
	return false
}
