package piscine

func IterativePower(nb int, power int) int {
	var result int
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := 0; i < power; i++ {
			result = result + nb
		}
		return result
	}
}
