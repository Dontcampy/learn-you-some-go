package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "French" {
		return "Bonjour, " + name
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
