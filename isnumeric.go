package piscine

func IsNumeric(str string) bool {
	var lengthOfString int
	var numberAmount int
	for _, item := range str {
		if item > 47 && item < 58 {
			numberAmount++
		}
		lengthOfString++
	}
	if numberAmount == lengthOfString {
		return true
	} else {
		return false
	}

}
