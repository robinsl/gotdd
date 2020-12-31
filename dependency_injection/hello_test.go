package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Christ")

	got := buffer.String()
	want := "Hello, Christ"

	if got != want {
		t.Errorf("got % q want %q", got, want)
	}
}
