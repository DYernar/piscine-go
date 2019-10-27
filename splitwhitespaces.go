package piscine

func SplitWhiteSpaces(str string) []string {
	var length int
	for range str {
		length++
	}
	arr := make([]string, length)
	var indWords string
	i := 0
	for _, item := range str {
		if item != ' ' && item != '\n' {
			indWords = indWords + string(item)
		} else {
			arr[i] = indWords
			i++
			indWords = ""
		}
	}
	return arr
}
