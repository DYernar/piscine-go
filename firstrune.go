package piscine

func FirstRune(s string) rune {
	var arr rune
	for index, item := range s {
		if index == 0 {
			arr = item
		}
	}
	return arr
}
