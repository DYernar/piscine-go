package piscine

func MakeRange(min, max int) []int {
	arr := make([]int, max-min)
	for i := min; i < max; i++ {
		arr[i] = i
	}
	return arr
}
