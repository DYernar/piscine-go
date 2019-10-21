package piscine

func Sqrt(nb int) int {
	var temp int
	if nb == 1 {
		return 1
	} else {
		for i := 0; i < nb; i++ {
			temp = i * i
			if temp == nb {
				return i
			}
		}
	}
	return 0
}
