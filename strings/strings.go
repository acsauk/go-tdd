package strings

import "strings"

func Uppercase(stringToUppercase string, retainWhiteSpace bool) string {
	uppercased := stringToUppercase
	if !retainWhiteSpace {
		strings.Trim(stringToUppercase, " ")
	}

	return strings.ToTitle(uppercased)
}