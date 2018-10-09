package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

func RepeatWithNoInitialization(count int) int {
	i := 0

	for i < count {
		i++
	}

	return i
}
