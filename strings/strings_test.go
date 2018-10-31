package strings

import "testing"

func TestUppercase(t *testing.T) {
	uppercased := Uppercase("hello there")
	expected := "HELLO THERE"

	if uppercased != expected {
		t.Errorf("expected '%s' but got '%s'", expected, uppercased)
	}
}