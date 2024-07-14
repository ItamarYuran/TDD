package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Itamar")
	want := "Hello, Itamar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
