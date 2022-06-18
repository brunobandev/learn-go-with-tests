package iteration

func Repeat(character string, loops int) string {
	var repeated string
	for i := 0; i < loops; i++ {
		repeated += character
	}

	return repeated
}
