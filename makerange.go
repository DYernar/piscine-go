package piscine

func MakeRange(min, max int) []int {
	arr := make([]int, max-min)

	if !(min >= max) {
		for i := min; i < max; i++ {
			arr[i] = i
		}
	}
	return arr

}
