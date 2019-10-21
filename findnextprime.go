package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	} else {
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
}
