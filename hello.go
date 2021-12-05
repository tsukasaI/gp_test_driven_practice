package main

import "fmt"

const englishHellloPrefix = "Hello, "
const spanishHellloPrefix = "Hola, "
const frenchHellloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHellloPrefix
	case "Spanish":
		prefix = spanishHellloPrefix
	default:
		prefix = englishHellloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
