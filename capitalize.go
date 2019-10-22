package piscine

func Capitalize(s string) string {
	var arr string
	var nextUpper bool
	for index, item := range s {
		if (item > 47 && item < 58) || (item > 64 && item < 91) || (item > 96 && item < 123) {
			nextUpper = false
		} else {
			nextUpper = true
		}
		if index == 0 {
			nextUpper = true
		}

		if nextUpper {
			arr = arr + ToUpper(string(item))
			nextUpper = false
		} else {
			arr = arr + ToLower(string(item))
		}
		if item == ' ' || item == '+' || item == '.' || item == ':' || item == ';' {
			nextUpper = true
		}
	}
	return arr
}
