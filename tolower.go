package piscine

func ToLower(s string) string {
	var arr string
	for _, item := range s {
		if item >= 65 && item <= 90 {
			arr = arr + string(item+32)
		} else {
			arr = arr + string(item)
		}
	}
	return arr
}
