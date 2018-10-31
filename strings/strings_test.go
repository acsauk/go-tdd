package strings

import "testing"
import "fmt"

func TestUppercaseAndTrim(t *testing.T) {
	uppercasedTrimmed := Uppercase("  hello there", true)
	expected := "HELLO THERE"

	if uppercasedTrimmed != expected {
		t.Errorf("expected '%s' but got '%s'", expected, uppercasedTrimmed)
	}
}

func TestUpperCaseAndDontTrim(t *testing.T) {
	uppercased := Uppercase("  hi there! ", false)
	expected := "  HI THERE! "

	if uppercased != expected {
		t.Errorf("expected '%s' but got '%s'", expected, uppercased)
	}
}

func ExampleUppercase() {
	uppercased := Uppercase(" hello there!!  ", true)
	fmt.Println(uppercased)
	// Output: HELLO THERE!!
}