package hello_world

import (
	"fmt"
)

const EnglishSentencePrefix = "Hello, "
const JapaneseSentencePrefix = "こんにちは, "
const FrenchSentencePrefix = "Bonjour, "
const PortugueseSentencePrefix = "Olá, "

func HelloWorld() string {
	return EnglishSentencePrefix + "World!"
}

func HelloWithName(name string) string {
	if name == "" {
		return EnglishSentencePrefix + "World!"
	}

	return fmt.Sprintf("%s%s!", EnglishSentencePrefix, name)
}

func PersonalisedHelloInLanguage(name, language string) string {
	languagePrefix := "Hello, "
	greeting := "World"

	switch language {
	case "French":
		languagePrefix = FrenchSentencePrefix
		greeting = "le Monde"
	case "English":
		languagePrefix = EnglishSentencePrefix
		greeting = "World"
	case "Portuguese":
		languagePrefix = PortugueseSentencePrefix
		greeting = "Mundo"
	case "Japanese":
		languagePrefix = JapaneseSentencePrefix
		greeting = "世界"
	}

	if name != "" {
		greeting = name
	}

	return fmt.Sprintf("%s%s", languagePrefix, greeting)
}
