package piscine

func Index(s string, toFind string) int {
	if StrLen(s) == 0 {
		return -1
	}
	if StrLen(toFind) == 0 {
		return 0
	}
	strrun := []rune(s)
	run := []rune(toFind)
	var res bool
	run1 := FirstRune(toFind)
	var result int
	for index, item := range s {
		if item == run1 {
			result = index
			res = true
			for i := range run {
				if (result + i) < StrLen(s) {
					if (StrLen(s) - result) >= StrLen(toFind) {
						if run[i] != strrun[result+i] {
							res = false
						}
					} else {
						return -1
					}

				} else {
					return -1
				}

			}
			if res == true {
				return result
			}
		}
	}
	if res == true {
		return result
	} else {
		return -1
	}

}
