package iteration

// Repeat repeats a character a number of reps times
func Repeat(character string, reps int) string {
	var repeated string
	for i := 0; i < reps; i++ {
		repeated += character
	}
	return repeated
}
