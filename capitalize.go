package piscine

func Capitalize(s string) string {
	var arr string
	var nextUpper bool
	for index, item := range s {
		if (item > 47 && item < 58) || (item > 64 && item < 91) || (item > 96 && item < 123) {
			if index == 0 {
				arr = arr + ToUpper(string(item))
			} else {
				if nextUpper {
					arr = arr + ToUpper(string(item))
					nextUpper = false
				} else {
					arr = arr + ToLower(string(item))
				}
			}
		} else {
			arr = arr + string(item)
			nextUpper = true
		}
	}
	return arr
}
