package main

import "fmt"

const helloEnPrefix = "Hello"
const helloFrPrefix = "Allo"
const helloSpPrefix = "Hola"
const french = "french"
const spanish = "spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", helloLanguagePrefix(language), name)
}

func helloLanguagePrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloFrPrefix
	case spanish:
		prefix = helloSpPrefix
	default:
		prefix = helloEnPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
