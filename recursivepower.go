package piscine

func RecursivePower(nb int, power int) int {
	var result int = nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return result * RecursivePower(nb, power-1)
	}
}
