package piscine

func IsPrime(nb int) bool {
	if nb < 0 {
		return false
	}
	if nb%2 != 0 || nb%3 != 0 || nb%5 != 0 || nb%7 != 0 || nb%13 != 0 {
		return true
	} else {
		return false
	}
}
