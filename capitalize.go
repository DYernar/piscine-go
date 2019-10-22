package piscine

func Capitalize(s string) string {
	var arr string
	var nextUpper bool
	for index, item := range s {
		if (item > 32 && item < 'A') || (item > 'Z' && item < 'a') || (item > 'z' && item < 128) {
			if nextUpper == true {
				nextUpper = true
			} else {
				nextUpper = false
			}
		}
		if index == 0 {
			nextUpper = true
		}

		if nextUpper && !(item > 32 && item < 'A') || (item > 'Z' && item < 'a') || (item > 'z' && item < 128) {
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
