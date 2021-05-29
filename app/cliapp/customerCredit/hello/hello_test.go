package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "hello world!!"
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
