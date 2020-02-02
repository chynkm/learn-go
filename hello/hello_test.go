package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("say hello to people", func(t *testing.T) {
        got := Hello("KM", "")
        want := "Hello KM"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello World"

        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("KM", "Spanish")
        want := "Hola KM"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("KM", "French")
        want := "Bonjour KM"
        assertCorrectMessage(t, got, want)
    })
}
