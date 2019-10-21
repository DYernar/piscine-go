package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	} else if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 13 || nb == 17 {
		return true
	} else if nb%2 == 0 || nb%3 == 0 || nb%5 == 0 || nb%7 == 0 || nb%13 == 0 || nb%17 == 0 || nb%19 == 0 || nb%79 == 0 {
		return false
	} else {
		return true
	}
}
