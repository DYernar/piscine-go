package piscine

func Join(strs []string, sep string) string {
	var length int
	var sum string
	for range strs {
		length++
	}
	for i := 0; i < length; i++ {
		sum = sum + sep + strs[i]
	}

	return sum
}
