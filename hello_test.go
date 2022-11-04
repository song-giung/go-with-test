package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("GiUng")
	want := "Hello, GiUng"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
