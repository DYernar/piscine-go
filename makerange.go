package piscine

func MakeRange(min, max int) []int {
	arr := make([]int, max-min)

	if !(min >= max) {
		for i := 0; i < max-min; i++ {
			arr[i] = i + min
		}
	}
	return arr

}
