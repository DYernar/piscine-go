package piscine

func IsNegative(nb int) rune {
	if nb < 0 {
		return 'F'
	} else if nb > 0 {
		return 'T'
	}else{
		return 0
	}
}