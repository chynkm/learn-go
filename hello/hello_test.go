package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("KM")
	want := "Hello KM"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
