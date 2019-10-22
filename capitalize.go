package piscine

func Capitalize(s string) string {
	var arr string
	for index, item := range s {
		if index == 0 {
			if (item < 48) || (item > 57 && item < 65) || (item > 90 && item < 97) || (item > 122) {
				return s
			}
			if item >= 97 && item <= 122 {
				arr = arr + string(item-32)
			} else {
				arr = arr + string(item)
			}
		} else {
			if item >= 65 && item <= 90 {
				arr = arr + string(item+32)
			} else {
				arr = arr + string(item)
			}
		}
	}
	return arr
}
