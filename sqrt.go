package piscine

func Sqrt(nb int) int {
	var temp int
	for i := 0; i < nb; i++ {
		temp = i * i
		if temp == nb {
			return i
		}
	}
	return 0
}
