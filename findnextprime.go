package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		for i := nb; ; i++ {
			if IsPrime(i) == true {
				return i
			}
		}
	}
}
