package deposit

// Calculate - рассчитывает параметры вклада
func Calculate(amount int, currency string) (min int, max int) {

	minPercent, maxPercent := PercentFor(currency)
	min = (amount * minPercent) / 1200
	max = (amount * maxPercent) / 1200
	return min, max
}

// PercentFor -
func PercentFor(currency string) (min int, max int) {
	if currency == "TJS" {
		return 4, 6
	} else if currency == "USD" {
		return 1, 2
	}
	return min, max
}
