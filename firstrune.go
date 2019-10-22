package piscine

func FirstRune(s string) rune {
	var arr []rune
	for index, item := range s {
		arr[index] = item
	}
	return arr[0]
}
