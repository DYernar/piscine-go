package piscine

func IsPrintable(str string) bool {
	var lengthOfString int
	var numberAmount int
	for _, item := range str {
		if item == '\n' {
			numberAmount++
		}
		lengthOfString++
	}
	if numberAmount > 0 {
		return false
	} else {
		return true
	}

}
