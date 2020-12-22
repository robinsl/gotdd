package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Robin", "")
		want := "Hello Robin!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world when there is no one to say hello to", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to a person in french", func(t *testing.T) {
		got := Hello("Robin", "french")
		want := "Allo Robin!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to a person in spanish", func(t *testing.T) {
		got := Hello("Robin", "spanish")
		want := "Hola Robin!"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloLanguagePrefix(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("return in english by default", func(t *testing.T) {
		got := helloLanguagePrefix("")
		want := "Hello"
		assertCorrectMessage(t, got, want)
	})

	t.Run("return in french when french is sent", func(t *testing.T) {
		got := helloLanguagePrefix("french")
		want := "Allo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("return in spanish when spanish is sent", func(t *testing.T) {
		got := helloLanguagePrefix("spanish")
		want := "Hola"
		assertCorrectMessage(t, got, want)
	})
}