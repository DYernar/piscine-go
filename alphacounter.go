package piscine

func AlphaCounter(str string) int {
	counter := 0
	for _, letter := range str {
		if (rune(letter) >= 65 && rune(letter) <= 90) || (rune(letter) >= 97 && rune(letter) <= 122) {
			counter++
		}
	}
	return counter
}
