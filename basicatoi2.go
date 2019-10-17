package piscine

func BasicAtoi2(s string) int {
	var i int
	for _, letter := range s {
		if letter == '1' {
			i = (i * 10) + 1
		} else if letter == '2' {
			i = (i * 10) + 2
		} else if letter == '3' {
			i = (i * 10) + 3
		} else if letter == '4' {
			i = (i * 10) + 4
		} else if letter == '5' {
			i = (i * 10) + 5
		} else if letter == '6' {
			i = (i * 10) + 6
		} else if letter == '7' {
			i = (i * 10) + 7
		} else if letter == '8' {
			i = (i * 10) + 8
		} else if letter == '9' {
			i = (i * 10) + 9
		} else if letter == '0' {
			i = i * 10
		} else if letter == ' ' {
			return 0
		} else {
			return 0
		}
	}
	return i
}
