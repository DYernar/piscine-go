package piscine

func Index(s string, toFind string) int {
	var length int
	var lengthToFind int
	foundIndex := 0
	for range toFind {
		lengthToFind++
	}
	for range s {
		length++
	}

	for i := 0; i < lengthToFind; i++ {
		for j := 0; j < length; j++ {
			if toFind[i] == s[j] {
				foundIndex = j
				return foundIndex
			}
		}
	}
	return 0
}
