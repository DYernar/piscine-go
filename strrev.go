package piscine

func StrRev(s string) string {
	var strVal string
	var i int
	for range s {
		i++
	}
	for ; i >=0; i-- {
		strVal += string (s[i])
	}
	return strVal
}