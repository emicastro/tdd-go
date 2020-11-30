package iteration

// Repeat receives a character and an n integer and return a string of character repeated n times
func Repeat(char string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}
