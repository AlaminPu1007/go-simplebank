package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	BDT = "BDT"
)

// isSupportedCurrency return true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR, BDT:
		return true
	}

	return false
}
