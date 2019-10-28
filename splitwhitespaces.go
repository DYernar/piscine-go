package piscine

func SplitWhiteSpaces(str string) []string {
	spaces := 0
	totalLength := 0
	for range str {
		totalLength++
	}
	for _, item := range str {
		if item == ' ' {
			spaces++
		}
	}
	whiteSpaces := make([]int, spaces+2)
	i := 0
	for index, item := range str {
		if index == 0 {
			whiteSpaces[i] = index
			i++
		}
		if item == ' ' || item == '\n' {
			whiteSpaces[i] = index
			i++
		}
	}
	whiteSpaces[i] = totalLength
	lengthOfArray := 0
	for range whiteSpaces {
		lengthOfArray++
	}
	arr := make([]string, spaces+1)
	for j := 0; j < lengthOfArray-1; j++ {
		if j == 0 {
			for k := whiteSpaces[j]; k < whiteSpaces[j+1]; k++ {
				arr[j] = arr[j] + string(str[k])
			}
		} else {
			if whiteSpaces[j+1]-whiteSpaces[j] != 1 {
				for k := whiteSpaces[j] + 1; k < whiteSpaces[j+1]; k++ {
					arr[j] = arr[j] + string(str[k])
				}
			}
		}

	}

	return arr
}
