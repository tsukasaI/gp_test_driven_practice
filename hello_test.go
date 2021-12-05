package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tsukasa")
	want := "Hello, Tsukasa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
