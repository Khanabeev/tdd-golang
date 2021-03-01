package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"
const russian = "Russian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Привет, "

func main() {
	fmt.Println(Hello("Anton", ""))
}

func Hello(s string, lang string) string {
	if s == "" {
		s = "World!"
	}
	return greetingPrefix(lang) + s
}

func greetingPrefix(lang string) string {
	var prefix string

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case english:
		prefix = englishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
