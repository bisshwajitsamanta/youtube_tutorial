package main

import "testing"

func Testhello(t *testing.T) {
	got := hello("Bishwajit")
	want := "Hello World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
