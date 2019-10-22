package piscine

func NthRune(s string, n int) rune {
	var arr rune
	for index, item := range s {
		if index == n+1 {
			arr = item
		}
	}
	return arr
}
