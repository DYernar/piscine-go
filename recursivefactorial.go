package piscine

func RecursiveFactorial(nb int) int {
	var result int = nb
	if nb > 0 && nb <= 25 {
		return result * RecursiveFactorial(nb-1)
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
