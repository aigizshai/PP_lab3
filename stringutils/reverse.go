package stringutils

func Reverse(str string) string {
	var result string
	for _, i := range str {
		result = string(i) + result
	}
	return result
}
