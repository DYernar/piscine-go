package piscine

func ToUpper(s string) string {
	var arr string
	for _, item := range s {
		if item >= 97 && item <= 122 {
			arr = arr + string(item-32)
		} else {
			arr = arr + string(item)
		}
	}
	return arr
}
