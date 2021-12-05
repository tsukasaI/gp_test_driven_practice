package main

import "fmt"

const englishHellloPrefix = "Hello, "
const spanishHellloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHellloPrefix + name
	}

	return englishHellloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
