package piscine

func IsAlpha(str string) bool {
	var lengthOfString int
	var lengthOfAlphas int
	for _, item := range str {
		if !((item > 47 && item < 58) || (item > 64 && item < 91) || (item > 96 && item < 123)) {
			lengthOfAlphas++
		}
		lengthOfString++
	}
	if lengthOfString == lengthOfAlphas {
		return false
	} else {
		return true
	}

}
