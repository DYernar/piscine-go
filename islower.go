package piscine

func IsNumeric(str string) bool {
	var lengthOfString int
	var lowerAmount int
	for _, item := range str {
		if item > 96 && item < 123 {
			lowerAmount++
		}
		lengthOfString++
	}
	if lowerAmount == lengthOfString {
		return true
	} else {
		return false
	}

}
