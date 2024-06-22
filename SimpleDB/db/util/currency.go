package util

const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
	IDR = "IDR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, IDR:
		return true
	}
	return false
}
