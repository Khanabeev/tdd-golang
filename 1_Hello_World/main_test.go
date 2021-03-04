package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Anton", "Spanish")
		want := "Hola, Anton"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mag", "French")
		want := "Bonjour, Mag"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Антон", "Russian")
		want := "Привет, Антон"
		assertCorrectMessage(t, got, want)
	})
}
