package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	got := DigitalRoot(16)
	want := 7

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		fmt.Println("success")
	}
}
