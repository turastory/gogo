package main

import "fmt"

const (
	helloPrefix       = "Hello "
	helloPrefixKorean = "안녕 "
	helloPrefixFrench = "Bonjour "
	korean            = "Korean"
	french            = "French"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case korean:
		prefix = helloPrefixKorean
	case french:
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
