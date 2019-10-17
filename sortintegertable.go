package piscine

func SortIntegerTable(table []int) []int {
	i := 0
	for range table {
		i++
	}
	for k := 0; k < i; k++ {
		for j := 0; j < i; j++ {
			if j != 0 {
				if table[j-1] > table[j] {
					temp := table[j-1]
					table[j-1] = table[j]
					table[j] = temp
				}
			}
		}
	}
	return table
}
