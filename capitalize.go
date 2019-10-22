package piscine

func Capitalize(s string) string {
	var arr string
	for index, item := range s {
		if index == 0 {
			arr = arr + ToUpper(string(item))
		} else {
			arr = arr + ToLower(string(item))
		}
	}
	return arr
}
