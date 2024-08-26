package iteration

func Repeat(c string, repeatCount int) string {
	var res string
	for i := 0; i < repeatCount; i++ {
		res += c
	}
	return res
}
