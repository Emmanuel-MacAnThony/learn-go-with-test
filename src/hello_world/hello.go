package helloworld

//import "fmt"

const englishHelloPrefix = "Hello, "
const spanishPrefix = "Hola, "
const spanish = "Spanish"
const frenchPrefix = "Bonjour, "
const french = "French"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
