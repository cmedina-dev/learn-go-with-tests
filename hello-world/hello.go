package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreeting  = "Bonjour, "
)

func getTranslatedHello(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getTranslatedHello(language) + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
