package iteration

func Repeat(s string, n int) string {
	// return strings.Repeat(s, n)
	var repeat string
	for i := 0; i < n; i++ {
		repeat += s
	}
	return repeat
}
