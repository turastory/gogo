package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Chris", "English")
		want := "Hello Chris"

		assert(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"

		assert(t, got, want)
	})

	t.Run("in Korean", func(t *testing.T) {
		got := Hello("Chris", "Korean")
		want := "안녕 Chris"

		assert(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thomas", "French")
		want := "Bonjour Thomas"

		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
