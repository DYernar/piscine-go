package piscine

func BasicJoin(strs []string) string {
	var length int
	var sum string
	for range strs {
		length++
	}
	for i := 0; i < length; i++ {
		sum = sum + strs[i]
	}

	return sum
}
