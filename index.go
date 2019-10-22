package piscine

func Index(s string, toFind string) int {
	var length int
	var lengthToFind int
	match := 0
	foundIndex := 0
	for range toFind {
		lengthToFind++
	}
	for range s {
		length++
	}

	for i := 0; i < lengthToFind; i++ {
		for j := 0; j < length; j++ {
			if (i == 0 && match == 0) || (i != 0 && match != 0) {
				if toFind[i] == s[j] {
					foundIndex = j
					match++
					return foundIndex
				}
			}
		}
	}
	return -1
}
