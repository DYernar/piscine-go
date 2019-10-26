package piscine

func ConcatParams(str []string) string {
	var indString string
	for index, item := range str {
		if index == 0 {
			indString = indString + item
		} else {
			indString = indString + "\n" + item
		}
	}
	return indString
}
