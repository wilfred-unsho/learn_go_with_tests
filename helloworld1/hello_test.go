package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCoreectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCoreectMessage(t, got, want)
	})
}

func assertCoreectMessage(t testing.TB, got, want string) {
	t.Helper()

    if got!= want {
        t.Errorf("got %v, want %v", got, want)
    }
}