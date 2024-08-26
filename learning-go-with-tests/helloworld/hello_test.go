package helloworld

import "testing"

// Think of t (*testing.T) as the hook into the testing framework,
// so that you can do things like t.Fail() when you want to fail.
func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello world when name is not provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello in French", func(t *testing.T) {
		got := Hello("Mr. Wick", "French")
		want := "Bonjour, Mr. Wick"
		assertCorrectMessage(t, got, want)
	})
	t.Run("speaks Bangla", func(t *testing.T) {
		got := Hello("Swagnik babu", "Bangla")
		want := "Nomoshkar, Swagnik babu"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB, which is an
// interface satisfied by both *testing.T and *testing.B so that you can call helper functions
// from a test or a benchmark.
//
// t.Helper() is needed to tell the test suite that this method is a helper. By doing this, when
// the test fails, the line number reported will be in our function call rather than inside our
// test helper.
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
