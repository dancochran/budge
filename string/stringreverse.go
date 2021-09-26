package string

// StringReverse takes a string and reverses it.
func StringReverse(str string) string {
	reversedString := make([]rune, len(str))
	for index := range str {
		reversedString[index] = rune(str[len(str)-1-index])
	}
	return string(reversedString)
}
