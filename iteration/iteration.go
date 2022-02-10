package iteration

const numberOfRepetitions = 5

func Repeat(char string) string {
	var repetitions string

	for i := 0; i < numberOfRepetitions; i++ {
		repetitions += char
	}

	return repetitions
}
