package iteration

func Repeat(s string)  string  {

	var repeated string
	for i :=0; i < 3 ; i++ {
		repeated += s
	}

	return repeated
}