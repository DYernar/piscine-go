package piscine

func SplitWhiteSpaces(str string) []string {
	bn := 1
	var invi1 rune
	for index, w := range str {
		if index == 0 {
			continue
		}
		if w == ' ' || w == '\t' || w == '\n' {
			if invi1 != ' ' && invi1 != '\t' && invi1 != '\n' {
				bn++
			}
		} else {

		}
		invi1 = w
	}

	answer := make([]string, bn)
	cnt := 0
	var invi rune
	for index, w := range str {
		if index == 0 {
			if w == ' ' || w == '\t' || w == '\n' {
				continue
			}
		}
		if w == ' ' || w == '\t' || w == '\n' {
			if invi != ' ' && invi != '\t' && invi != '\n' {
				cnt++
			}

		} else {

			answer[cnt] += string(w)

		}
		invi = w
	}

	return answer

}
