package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Friend", "")
		want := "Hello, Friend"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Helena", "Spanish")
		want := "Hola, Helena"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Dartagnan", "french")
		want := "Bonjour, Dartagnan"
		assertCorrectMessage(t, got, want)
	})
}
