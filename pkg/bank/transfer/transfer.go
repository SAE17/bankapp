package transfer

// Total is
func Total(amount int) (total int) {
	const percent = 5
	if amount >= 500000 {
		total = (amount*percent)/1000 + amount
	} else {
		total = amount
	}
	return total
}
