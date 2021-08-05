package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, GoKit!"
	got := SayHello()
	if got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	} else {
		t.Logf("SayHello() = %q.", got)
	}
}
