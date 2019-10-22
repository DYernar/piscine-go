package piscine

func IsAlpha(str string) bool {
	var lengthOfString int
	var notAlpha int
	for _, item := range str {
		if (item > 47 && item < 58) || (item > 64 && item < 91) || (item > 96 && item < 123) {
			notAlpha++
		}
		lengthOfString++
	}
	if notAlpha > 0 {
		return false
	} else {
		return true
	}

}
