package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello world!"
	if got := main(); got != want {
		t.Errorf("main() = %q, want =%q", got, want)
	}

}
