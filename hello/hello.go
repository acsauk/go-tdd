package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"

const helloPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "
const helloGermanPrefix = "Halo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloFrenchPrefix
	case spanish:
		prefix = helloSpanishPrefix
	case german:
		prefix = helloGermanPrefix
	default:
		prefix = helloPrefix
	}

	return 
}

func main() {
	fmt.Println(Hello("world", ""))
}
