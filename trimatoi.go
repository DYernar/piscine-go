package piscine

func TrimAtoi(str string) int {
	var arr string
	for _, item := range str {
		if item == '-' && len(arr) == 0 {
			arr = string(item) + arr
		}
		if item >= 48 && item <= 57 {
			arr = arr + string(item)
		}
	}
	return Atoi(arr)
}
