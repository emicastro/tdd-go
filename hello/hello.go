package main

import (
	"fmt"
	"strings"
)

const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const englishGreetingPrefix = "Hello, "
const defaultGreetingPostfix = "World"
const spanish = "spanish"
const french = "french"

// Hello returns a greeting string
func Hello(name, language string) string {

	if name == "" {
		name = defaultGreetingPostfix
	}

	return greetingPrefix(strings.ToLower(language)) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World!", ""))
}
