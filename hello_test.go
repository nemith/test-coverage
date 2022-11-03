package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World"
	got := Hello()

	if want != got {
		t.Errorf("unexpected results (got: %q, want: %q)", got, want)
	}

}
