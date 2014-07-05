package leap

func IsLeapYear(year int) bool {
	if !isDivisibleBy(year, 4) {
		return false
	}
	if !isDivisibleBy(year, 100) {
		return true
	}
	if !isDivisibleBy(year, 400) {
		return false
	}
	return true
}

func isDivisibleBy(year, divisor int) bool {
	return year%divisor == 0
}
