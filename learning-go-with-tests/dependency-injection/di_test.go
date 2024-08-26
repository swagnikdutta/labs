package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Chris")

	got := buf.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
