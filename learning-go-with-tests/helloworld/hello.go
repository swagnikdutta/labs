package helloworld

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	bangla  = "Bangla"

	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	banglaHelloPrefix  = "Nomoshkar, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case bangla:
		prefix = banglaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
