package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, welcome to gopher"
	received := Hello()

	if expected != received {
		t.Fatalf("%s was expected but %s arrived", expected, received)
	}
}
