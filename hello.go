package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

const englishPrefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishPrefixHello + name + "!"
}
