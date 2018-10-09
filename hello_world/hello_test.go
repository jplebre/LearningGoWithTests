package hello_world

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestHello(t *testing.T) {
	got := HelloWorld()
	expected := "Hello, World!"

	if got != expected {
		t.Errorf("got '%s' want '%s'", got, expected)
	}
}

func TestHelloNestedTests(t *testing.T) {
	t.Run("Program says hello to people", func(t *testing.T) {
		got := HelloWithName("Chris")
		expected := "Hello, Chris!"

		if got != expected {
			t.Errorf("got '%s' want '%s'", got, expected)
		}
	})

	t.Run("Program says Hello, World if the name is empty", func(t *testing.T) {
		got := HelloWithName("")
		expected := "Hello, World!"

		if got != expected {
			t.Errorf("got '%s' want '%s'", got, expected)
		}
	})

}

func TestHelloWithGomega(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(HelloWorld()).To(Equal("Hello, World!"))
}

func TestHelloWithGomegaGlobally(t *testing.T) {
	RegisterTestingT(t) // This binds T globally, which means tests can't run in parallel
	Expect(HelloWorld()).To(Equal("Hello, World!"))
}

func TestHelloWithArguments(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(HelloWithName("Chris")).To(Equal("Hello, Chris!"))
}

func TestHelloWithLanguages(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper() // Tells the test suite to report the error at the test level, not at this helper level
		g := NewGomegaWithT(t)

		g.Expect(got).To(Equal(expected))
	}

	t.Run("Program Says Hello In In French", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("", "French"), "Bonjour, le Monde")
	})

	t.Run("Program Says Hello In Portuguese", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("", "Portuguese"), "Olá, Mundo")
	})

	t.Run("Program Says Hello In Japanese", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("", "Japanese"), "こんにちは, 世界")
	})

	t.Run("Program Returns Personalised greeting In French", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("Chris", "French"), "Bonjour, Chris")
	})

	t.Run("Program Says Hello In Portuguese", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("Chris", "Portuguese"), "Olá, Chris")
	})

	t.Run("Program Says Hello In Japanese", func(t *testing.T) {
		assertCorrectMessage(t, PersonalisedHelloInLanguage("Chris", "Japanese"), "こんにちは, Chris")
	})
}
