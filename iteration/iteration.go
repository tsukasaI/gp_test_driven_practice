package iteration

func Repeat(str string, time int) string {
	var result string
	for i := 0; i < time; i++ {
		result += str
	}
	return result
}
