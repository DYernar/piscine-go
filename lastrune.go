package piscine

func LastRune(s string) rune {
	var length int
	var arr []rune = []rune(s)
	for range s {
		length++
	}
	return arr[length-1]
}
