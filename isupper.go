package piscine

func IsUpper(str string) bool {
	var lengthOfString int
	var upperAmount int
	for _, item := range str {
		if item > 64 && item < 91 {
			upperAmount++
		}
		lengthOfString++
	}
	if upperAmount == lengthOfString {
		return true
	} else {
		return false
	}

}
