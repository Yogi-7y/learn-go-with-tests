package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

const (
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}
