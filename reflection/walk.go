package walk

func walk(x interface{}, fn func(input string)) {
	fn("Test string")
}